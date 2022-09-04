package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func TransformString(str string) string {
	data := md5.Sum([]byte(str))
	return hex.EncodeToString(data[:])
}
