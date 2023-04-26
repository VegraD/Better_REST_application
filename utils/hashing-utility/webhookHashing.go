package hashing_utility

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashingTheWebhook(url string, country string, calls int) string {
	hash := hmac.New(sha256.New, getSecret())
	hash.Write([]byte(fmt.Sprintf("%s%s%d", url, country, calls)))
	return hex.EncodeToString(hash.Sum(nil))
}
