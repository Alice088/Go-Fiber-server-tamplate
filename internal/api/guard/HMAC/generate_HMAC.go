package HMAC

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func Generate(signature string) string {
	mac := hmac.New(sha256.New, []byte(os.Getenv("HMAC_KEY")))
	mac.Write([]byte(signature))
	return hex.EncodeToString(mac.Sum(nil))
}
