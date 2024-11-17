package HMAC

import (
	"crypto/hmac"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
	"star_trade/backend/internal/core/support"
	"strconv"
	"time"
)

const fiveMinute = 300

func ValidateHMAC(r *http.Request, session *sessions.Session) bool {
	if support.WasUsed(r.URL.Query().Get("signature"), &session.Values) {
		return false
	}

	timestampStr := r.URL.Query().Get("time")
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	timeOfLife := support.MakePositive(timestamp - time.Now().Unix())
	if err != nil {
		return false
	}

	if timeOfLife > fiveMinute {
		return false
	}

	message := fmt.Sprintf("%s%s", r.URL.Path, timestampStr)
	receivedMAC := r.URL.Query().Get("signature")
	return hmac.Equal([]byte(receivedMAC), []byte(GenerateHMAC(message)))
}
