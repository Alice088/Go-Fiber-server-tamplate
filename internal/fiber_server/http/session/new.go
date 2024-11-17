package session

import (
	"github.com/gorilla/sessions"
)

func New() *sessions.FilesystemStore {
	store := sessions.NewFilesystemStore(
		"",
		[]byte("707e4b95025a528341deadb3a1"),
	)

	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   60 * 6,
		HttpOnly: true,
	}

	return store
}
