package utility

import (
	"crypto/sha1"
	"fmt"
)

func ShaHash(str string) (string, error) {
	passSha1 := sha1.New()
	_, err := passSha1.Write([]byte(str))
	if err != nil {
		return str, err
	}

	getSha1 := passSha1.Sum(nil)
	return fmt.Sprintf("%x", getSha1), nil
}
