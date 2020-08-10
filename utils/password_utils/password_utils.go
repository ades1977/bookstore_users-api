package password_utils

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	_ "github.com/ades1977/bookstore_users-api/utils/errors"
)

const (
	// Size is the size, in bytes, of a SHA-512 checksum.
	Size = 64
	// Size224 is the size, in bytes, of a SHA-512/224 checksum.
	Size224 = 28
	// Size256 is the size, in bytes, of a SHA-512/256 checksum.
	Size256 = 32
	// Size384 is the size, in bytes, of a SHA-384 checksum.
	Size384 = 48
	// BlockSize is the block size, in bytes, of the SHA-512/224,
	// SHA-512/256, SHA-384 and SHA-512 hash functions.
	BlockSize = 128
	ipassword= "hello"
)


func GetMD5(ipass string) string{
	hash := md5.New()
	hash.Write([]byte(ipass))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetSha(ipass string,xsize  string) string{
	hash := sha512.New()
	hash.Sum([]byte(xsize))
	hash.Write([]byte(ipass))
	return hex.EncodeToString(hash.Sum(nil))
}

