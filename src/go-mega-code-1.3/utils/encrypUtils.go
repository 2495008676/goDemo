package utils1

import (
	"crypto/md5"
	"encoding/hex"
)

/**
加密统一入口
*/
func GeneratePasswordHash(pwd string) string {
	return Md5(pwd)
}

/**
md5加密方法
*/
func Md5(origin string) string {
	hasher := md5.New()
	hasher.Write([]byte(origin))
	return hex.EncodeToString(hasher.Sum(nil))
}
