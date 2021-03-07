package utils

import (
	"encoding/base64"
)

func Base64Encode(origin string) (result string) {
	result = base64.StdEncoding.EncodeToString([]byte(origin))
	return result
}

func Base64Decode(origin string) (result string) {
	tmp, _ := base64.StdEncoding.DecodeString(origin)
	return string(tmp)
}
