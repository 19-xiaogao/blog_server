package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

var letters = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func CreateSixNumber() string {
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
