package api_key

import (
	"crypto/sha256"
	"crypto/subtle"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func Validate(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}
