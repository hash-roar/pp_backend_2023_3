package utils

import (
	"crypto/md5"
	"fmt"
)

func IsEqual(encryptStr string, toEncryptStr string) bool {
	encryptedStr := fmt.Sprintf("%x", md5.Sum([]byte(toEncryptStr)))
	return encryptStr == encryptedStr
}

func Md5(toEncryptStr string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(toEncryptStr)))
}
