package main

import (
	"crypto/md5"
	"encoding/base64"
)

func ToMd5String(str string) string {
	temp := md5.Sum([]byte(str))
	return base64.StdEncoding.EncodeToString(temp[:])
}

func Bool2Type(b bool) uint {
	if b {
		return IS_DIR
	}
	return IS_FILE
}

func DecodeB64String(s string) string {
	res, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return ""
	}
	return string(res)
}

const (
	IS_PRO = true
)
