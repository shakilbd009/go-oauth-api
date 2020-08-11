package crypto_utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(pass string) string {

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(pass))
	return hex.EncodeToString(hash.Sum(nil))
}
