package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/sessions"
)

type CreateSessionRequest struct {
	W        http.ResponseWriter
	R        *http.Request
	Username string
}


func CreateSession(data CreateSessionRequest) (*sessions.Session, error) {

	var store = sessions.NewFilesystemStore("tmp/sessions", []byte(os.Getenv("SESSION_SECRET")))
	maxAge, err := strconv.Atoi(os.Getenv("SESSION_MAX_AGE"))
	if err != nil {
		maxAge = 3600
	}

	session, _ := store.Get(data.R, "session")
	session.Values["username"] = data.Username
	session.Options.MaxAge = maxAge

	err = session.Save(data.R, data.W)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func generateSessionID() (string, error) {
	// Generate 16 random bytes
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to a hex string
	sessionID := hex.EncodeToString(randomBytes)
	return sessionID, nil
}

func GetSession(r *http.Request, sessionID string) (*sessions.Session, error) {
	var store = sessions.NewFilesystemStore("tmp/sessions", []byte(os.Getenv("SESSION_SECRET")))

	session, err := store.Get(r, sessionID)
	if err != nil {
		return nil, errors.New("Session not found")
	}

	return session, nil
}
