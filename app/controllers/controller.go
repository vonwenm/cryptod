package controllers

import (
	r "github.com/revel/revel"
)

const (
	RespCodeTag = "resp_code"
	ErrorsTag   = "errors"
	RemarkTag   = "remark"
	OutputTag   = "output"

	ResponseSuccess = 1000
	ResponseError   = 9000

	RemarkSuccess = "success"
	RemarkError   = "error"
)

const (
	MethodMD4 = iota
	MethodMD5
	MethodSHA1
	MethodSHA256
	MethodSHA384
	MethodSHA512

	CodecRaw = iota
	CodecHex
	CodecBase32
	CodecBase64

	CodecRawStr = "raw"
	CodecRawB32 = "b32"
	CodecRawB64 = "b64"
	CodecRawHex = "hex"

	ModeEnc = "enc"
	ModeDec = "dec"
)

type BaseController struct {
	*r.Controller
	jsonOutput map[string]interface{}
	errors     []interface{}
}

func (c *BaseController) initJsonOutput() {
	if c.jsonOutput == nil {
		c.jsonOutput = make(map[string]interface{})
	}
}

//Error 设置失败操作标志
func (c *BaseController) Error(err interface{}) {
	c.Response(RespCodeTag, ResponseError)
	c.Response(RemarkTag, RemarkError)
	c.errors = append(c.errors, err)
}

// Success 设置成功操作标志
func (c *BaseController) Success() {
	c.Response(RespCodeTag, ResponseSuccess)
	c.Response(RemarkTag, RemarkSuccess)
}

// Response 输出内容赋值
func (c *BaseController) Response(key string, val interface{}) *BaseController {
	c.initJsonOutput()
	c.jsonOutput[key] = val
	return c
}

// RenderAPIJSON 渲染 api json 输出
func (c *BaseController) RenderAPIJSON() r.Result {
	c.initJsonOutput()
	if _, ok := c.jsonOutput[RespCodeTag]; !ok {
		c.Response(RespCodeTag, ResponseSuccess)
	}
	if len(c.errors) > 0 {
		c.Response(RespCodeTag, ResponseError)
		c.Response(ErrorsTag, c.errors)
	}
	return c.RenderJson(c.jsonOutput)
}
