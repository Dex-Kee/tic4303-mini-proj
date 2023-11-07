package util

import (
	"math/rand"
	"regexp"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	emailPattern          = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneNumberPattern    = regexp.MustCompile(`^[6789]\d{7}$`)
	strongPasswordPattern = regexp.MustCompile(`^.*[A-Z].*[a-z].*\d.*[@#$%^&+=].*$`)
	seededRand            = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func IsEmail(str string) bool {
	return emailPattern.MatchString(str)
}

func IsPhoneNumber(str string) bool {
	return phoneNumberPattern.MatchString(str)
}

func IsStrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	return strongPasswordPattern.MatchString(password)
}
