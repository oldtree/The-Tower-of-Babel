package utils

import (
	//"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	//"crypto/sha256"
	"encoding/hex"
	"fmt"
	//"hash"
)

func Debug(infos ...interface{}) {
	if true {
		fmt.Printf("DEBUG: "+fmt.Sprintf("%s\n", infos[0]), infos[1:]...)
	}
}

func GetRandomString(n int) string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

// Encode string to md5 hex value
func EncodeMd5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
