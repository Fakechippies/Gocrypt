package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

func Parser(hashMethod string, hashString string) string {
	if hashMethod == "md5" {
		return md5Conv(hashString)
	}
	return "Unsupported Hash Conversion"
}

func md5Conv(hashString string) string {
	data := []byte(hashString)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
