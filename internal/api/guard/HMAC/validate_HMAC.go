package HMAC

import (
	"app/internal/api/custom_errors"
	"crypto/hmac"
	"fmt"
	"github.com/gorilla/sessions"
	"math"
	"net/http"
	"strconv"
	"time"
)

const twoMinute = int64(2 * time.Minute)

func Validate(r *http.Request, session *sessions.Session) (bool, error) {
	URLSignature := r.URL.Query().Get("signature")
	URLTime := r.URL.Query().Get("time")

	if IsRepeated(URLSignature, &session.Values) {
		return false, custom_errors.ErrRepeatedRequest
	}

	timestamp, err := strconv.ParseInt(URLTime, 10, 64)
	timeOfLife := int64(math.Abs(float64(timestamp - time.Now().Unix())))
	if err != nil {
		return false, err
	}

	if timeOfLife > twoMinute {
		return false, custom_errors.ErrKeyExpired
	}

	NewSignature := fmt.Sprintf("%s%s", r.URL.Path, URLTime)
	return hmac.Equal([]byte(URLSignature), []byte(Generate(NewSignature))), nil
}
