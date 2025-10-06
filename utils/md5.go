package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func Md5Encode(plainText string) string {
	h := md5.New()
	h.Write([]byte(plainText))
	tempStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(tempStr))
}

func GenerateSalt(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func MakePassword(plainpwd, salt string) string {
	return Md5Encode(plainpwd + salt)
}

func ValidPassword(plainpwd, salt, password string) bool {
	return Md5Encode(plainpwd+salt) == password
}
