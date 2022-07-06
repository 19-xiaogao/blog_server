package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

var letters = []rune("1A2B3C4D5E6F7G8H9IJKLMNOPQRSTUVWXYZ")

func CreateSixNumber() string {
	b := make([]rune, 6)
	rand.Seed(time.Now().UnixNano())
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
