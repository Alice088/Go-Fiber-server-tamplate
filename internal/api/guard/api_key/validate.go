package api_key

import (
	"crypto/hmac"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"log/slog"
)

func Validate(key string, logger *slog.Logger) (bool, error) {
	hashedAPIKey, err := Generate(logger)
	if err != nil {
		return false, err
	}

	if hmac.Equal([]byte(hashedAPIKey), []byte(key)) {
		return true, nil
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}
