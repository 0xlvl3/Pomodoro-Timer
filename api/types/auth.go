package types

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"time"
)

type Session struct {
	Token     string
	UserID    string
	CreatedAt time.Time
	ExpiresAt time.Time
}

func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

var sessions = make(map[string]Session)

func loginUser(username, password string) string {
	// verify user

	// create session token
	token := GenerateRandomString(32)

	// store session token in your session store

	// return session token to client
	return token
}

func authenticateUser(token string) bool {
	// verify the session token
	session, exists := sessions[token]
	if exists && time.Now().Before(session.ExpiresAt) {
		return true
	}

	//if the session is not valid, delete it from the session store
	delete(sessions, token)
	return false
}
