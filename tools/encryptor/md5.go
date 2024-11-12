package encryptor

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
func CheckMd5(src []byte, md5str string) bool {
	return Md5(src) == md5str
}
