package utils

import (
	"golang.org/x/exp/rand"
	"strconv"
)

func GenerateOtpCode() string {
	return strconv.Itoa(rand.Intn(999999-99999) + 99999)
}
