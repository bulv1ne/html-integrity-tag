package htmlintegrity

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

func CalculateSHA384(data []byte) []byte {
	hash := sha512.Sum384(data)
	return hash[:]
}

func CalculateSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func IntegirtyString(data []byte, shaSize int) string {
	switch shaSize {
	case 256:
		hash := CalculateSHA256(data)
		return "sha256-" + base64.StdEncoding.EncodeToString(hash)
	case 384:
		hash := CalculateSHA384(data)
		return "sha384-" + base64.StdEncoding.EncodeToString(hash)
	default:
		panic(fmt.Errorf("unsupported SHA size: %d", shaSize))
	}
}
