package controllers

// 数字摘要算法: md4,md5,sha1,sha256,sha384,sha512
// 对传入的参数值做数字摘要,支持 GET|POST 提交方式
// 接收参数
// raw: 需要做摘要的值
// incode: raw|hex|b64|b32 传入值的编码方式,默认 raw
// outcode: hex|b64|b32 输出的编码方式,默认 b64
// 响应
// {
//     resp_code: 1000,
//     errors: [如果有错误发生,列出错误描述],
//     output: "摘要输出"
// }
//
//

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"github.com/revel/revel"
	"golang.org/x/crypto/md4"
	"hash"
	"github.com/weidewang/cryptod/app"
)

type Digest struct {
	BaseController
}

// Index 输出帮助页面
func (d Digest) Index() revel.Result {
	return d.Render()
}

// 具体处理方法
func (d Digest) DoDigest(raw, incode, outcode, method string) revel.Result {

	if len(incode) == 0 {
		incode = CodecRawStr
	}
	if len(outcode) == 0 {
		outcode = CodecRawB64
	}

	codec := func(c string) int {
		switch c {
		case CodecRawHex:
			return CodecHex
		case CodecRawB32:
			return CodecBase32
		case CodecRawB64:
			return CodecBase64
		default:
			return CodecRaw
		}
	}

	var err error
	var rawBytes []byte
	var h hash.Hash

	switch codec(incode) {
	case CodecHex:
		rawBytes, err = app.HexToBytes(raw)
	case CodecBase64:
		rawBytes, err = app.Base64ToBytes(raw)
	default:
		rawBytes = []byte(raw)
	}

	if err != nil {
		d.Error(err.Error())
		d.Response(RemarkTag, RemarkError)
		return d.RenderAPIJSON()
	}

	d.Response("method", method)
	d.Response("mode", ModeEnc)

	methodCode := func(m string) int {
		switch m {
		case "md4":
			return MethodMD4
		case "md5":
			return MethodMD5
		case "sha1":
			return MethodSHA1
		case "sha256":
			return MethodSHA256
		case "sha384":
			return MethodSHA384
		case "sha512":
			return MethodSHA512
		}
		return MethodMD5
	}

	switch methodCode(method) {
	case MethodMD4:
		h = md4.New()
	case MethodMD5:
		h = md5.New()
	case MethodSHA1:
		h = sha1.New()
	case MethodSHA256:
		h = sha256.New()
	case MethodSHA384:
		h = sha512.New384()
	case MethodSHA512:
		h = sha512.New()
	}

	h.Write(rawBytes)
	switch codec(outcode) {
	case CodecBase32:
		d.Response(OutputTag, app.BytesToBase32(h.Sum(nil)))
	case CodecBase64:
		d.Response(OutputTag, app.BytesToBase64(h.Sum(nil)))
	default:
		d.Response(OutputTag, app.BytesToHex(h.Sum(nil)))
	}

	return d.RenderAPIJSON()
}
