package util

import (
	"math/rand"
	"regexp"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

const phoneNumberPattern = `^[6789]\d{7}$`

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func IsEmail(str string) bool {
	matched, _ := regexp.MatchString(emailPattern, str)
	return matched
}

func IsPhoneNumber(str string) bool {
	matched, _ := regexp.MatchString(phoneNumberPattern, str)
	return matched
}
