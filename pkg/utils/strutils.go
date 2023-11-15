package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/duke-git/lancet/v2/strutil"
)

func IsAnyStringBlank(strs ...string) bool {
	for _, str := range strs {
		if strutil.IsBlank(str) {
			return true
		}
	}
	return false
}

func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str)) // 需要加密的字符串为
	return hex.EncodeToString(h.Sum(nil))
}
