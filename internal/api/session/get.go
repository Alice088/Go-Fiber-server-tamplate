package session

import (
	"github.com/gorilla/sessions"
	"net/http"
)

func Get(sessionName string, r *http.Request, store *sessions.FilesystemStore) (*sessions.Session, error) {
	session, err := store.Get(r, sessionName)

	if err != nil {
		return nil, err
	}

	return session, nil
}
