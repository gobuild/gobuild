package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandNString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func StrFormat(format string, kv map[string]interface{}) string {
	for key, val := range kv {
		key = "{" + key + "}"
		format = strings.Replace(format, key, fmt.Sprintf("%v", val), -1)
	}
	return format
}
