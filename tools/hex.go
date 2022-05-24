package tools

import (
	"crypto/md5"
	"fmt"
)

const salt = "ux$ad70*b"

func Md5Hex(msg string) string {
	data := []byte(msg)
	hash := md5.Sum(data)
	md5str := fmt.Sprintf("%x", hash)

	// 二次加密
	fmt.Println(md5str + salt)
	hash = md5.Sum([]byte(md5str + salt))
	md5str = fmt.Sprintf("%x", hash)

	return md5str
}
