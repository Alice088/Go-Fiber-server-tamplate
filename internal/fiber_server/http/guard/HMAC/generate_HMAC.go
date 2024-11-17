package HMAC

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"go.deanishe.net/env"
)

func GenerateHMAC(message string) string {
	mac := hmac.New(sha256.New, []byte(env.Get("HMAC_KEY")))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
