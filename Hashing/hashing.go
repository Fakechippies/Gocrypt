package hashing

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Parser(hashMethod string, hashString string) string {
	if hashMethod == "md5" {
		return md5Conv(hashString)
	} else if hashMethod == "sha1" {
		return sha1Conv(hashString)
	} else if hashMethod == "sha256" {
		return sha256Conv(hashString)
	} else if hashMethod == "sha224" {
		return sha224Conv(hashString)
	} else if hashMethod == "sha384" {
		return sha384Conv(hashString)
	} else if hashMethod == "sha512" {
		return sha512Conv(hashString)
	}
	return "Unsupported Hash Conversion"
}

func md5Conv(hashString string) string {
	data := []byte(hashString)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func sha1Conv(hashString string) string {
	data := []byte(hashString)
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

func sha256Conv(hashString string) string {
	data := []byte(hashString)
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

func sha224Conv(hashString string) string {
	data := []byte(hashString)
	hash := sha256.Sum224(data)
	return hex.EncodeToString(hash[:])
}

func sha384Conv(hashString string) string {
	data := []byte(hashString)
	hash := sha512.Sum384(data)
	return hex.EncodeToString(hash[:])
}

func sha512Conv(hashString string) string {
	data := []byte(hashString)
	hash := sha512.Sum512(data)
	return hex.EncodeToString(hash[:])
}
