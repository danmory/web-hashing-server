package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func md5Hash(str string) string {
	data := md5.Sum([]byte(str))
	return hex.EncodeToString(data[:])
}

type converter[T any] interface {
	Do(v T) T
}

type stringConverter struct {
	convertFunc func(string) string
}

func (sc *stringConverter) Do(str string) string {
	return sc.convertFunc(str)
}

var StringConverter converter[string] = &stringConverter{md5Hash}
