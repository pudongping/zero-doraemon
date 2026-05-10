package hashx

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 生成字符串的 MD5 哈希值
func MD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
