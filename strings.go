package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5str(s string) string  {
	h := md5.New()
	h.Write([]byte(s))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
