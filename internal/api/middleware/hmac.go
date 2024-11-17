package middleware

import (
	"RuRu/internal/api/guard/HMAC"
	supportSession "RuRu/internal/api/session"
	"github.com/gorilla/sessions"
	"log/slog"
	"net/http"
)

func Hmac(logger *slog.Logger, store *sessions.FilesystemStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		logger = logger.With(slog.String("component", "custom_middleware/hmac"))
		logger.Info("hmac middleware enabled")

		fn := func(w http.ResponseWriter, r *http.Request) {
			session, err := supportSession.Get("authSave", r, store)

			if err != nil {
				logger.Error(err.Error())
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}

			ok, err := HMAC.Validate(r, session)

			err = session.Save(r, w)
			if err != nil {
				logger.Error(err.Error())
				http.Error(w, "Something went wrong", http.StatusInternalServerError)
				return
			}

			if !ok {
				logger.Error("Unauthorized request")
				http.Error(w, "Unauthorized request", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
