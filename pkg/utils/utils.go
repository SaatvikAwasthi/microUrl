package utils

import (
	"encoding/base64"
)

func Base64Encode(data string) string {
	encoder := base64.RawURLEncoding.WithPadding(base64.NoPadding)
	return encoder.EncodeToString([]byte(data))
}

func Base64EncodeWithPadding(data string, pad rune) string {
	encoder := base64.RawURLEncoding.WithPadding(pad)
	return encoder.EncodeToString([]byte(data))
}

func CheckEmpty(data string) bool {
	return data == ""
}
