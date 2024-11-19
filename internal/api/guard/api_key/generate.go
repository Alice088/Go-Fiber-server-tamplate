package api_key

import (
	"RuRu/internal/api/custom_errors"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"os"
)

func Generate(logger *slog.Logger) (string, error) {
	var apiKey string
	var apiKeySalt string

	if apiKey = os.Getenv("API_KEY"); len(apiKey) == 0 {
		logger.Error("API_KEY in .env doesn't set")
		return "", custom_errors.ErrServerSideError
	}

	if apiKeySalt = os.Getenv("API_KEY_SALT"); len(apiKeySalt) == 0 {
		logger.Error("API_KEY_SALT in .env doesn't set")
		return "", custom_errors.ErrServerSideError
	}

	mac := hmac.New(sha256.New, []byte(apiKey))
	mac.Write([]byte(apiKeySalt))
	return hex.EncodeToString(mac.Sum(nil)), nil
}
