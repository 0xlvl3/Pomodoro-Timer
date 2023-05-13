package handles

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func generateTokenWithExpiry(userID string) (string, error) {
	// The JWT key used to sign tokens
	var jwtKey = []byte("your_secret_key")

	expirationTime := time.Now().Add(2 * time.Hour) // Token valid for 5 minutes

	// Create the JWT claims, which includes the username and expiry time
	claims := &jwt.StandardClaims{
		// In JWT, the expiry time is expressed as unix milliseconds
		ExpiresAt: expirationTime.Unix(),
		Issuer:    userID,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}

	return tokenString, nil
}

func validateToken(tknStr string) (bool, error) {
	var jwtKey = []byte("your_secret_key")

	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, nil
		}
		return false, err
	}
	if !tkn.Valid {
		return false, nil
	}
	if claims.ExpiresAt < time.Now().Unix() {
		return false, errors.New("Token expired")
	}

	return true, nil
}
