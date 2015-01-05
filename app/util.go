package app

import (
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
)

// HexToBytes 还原 hex 编码值为 byte 数组
func HexToBytes(h string) ([]byte, error) {
	return hex.DecodeString(h)
}

// BytesToHex 将 bytes 编码成 hex string
func BytesToHex(bs []byte) string {
	return hex.EncodeToString(bs)
}

// Base64ToBytes 还原 base64 编码值为 byte 数组,标准 base64
func Base64ToBytes(bs string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(bs)
}

// BytesToBase64  将 bytes 编码成 base64 字符串
func BytesToBase64(bs []byte) string {
	return base64.StdEncoding.EncodeToString(bs)
}

// Base32ToBytes 还原 base32 编码值为 byte 数组,标准 base32
func Base32ToBytes(bs string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(bs)
}

// BytesToBase32  将 bytes 编码成 base32 字符串
func BytesToBase32(bs []byte) string {
	return base32.StdEncoding.EncodeToString(bs)
}
