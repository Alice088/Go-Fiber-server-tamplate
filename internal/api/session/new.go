package session

import (
	"github.com/gorilla/sessions"
)

func New() *sessions.FilesystemStore {
	store := sessions.NewFilesystemStore(
		"",
		[]byte("<BYTES>"),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 6,
		HttpOnly: true,
	}

	return store
}
