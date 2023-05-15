package handles

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/0xlvl3/pomodoro-timer/db"
	"github.com/0xlvl3/pomodoro-timer/types"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func generateTokenWithExpiry(userStore db.UserStore) (valid bool) {

	// The JWT key used to sign tokens
	var jwtKey = []byte("your_secret_key")

	expirationTime := time.Now().Add(2 * time.Hour) // Token valid for 5 minutes

	// Create the JWT claims, which includes the username and expiry time
	claims := validateToken(token)

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

// creating our token
// TODO: make this private again. -> we made it public in our seed.
func CreateTokenFromUser(user *types.User) string {
	now := time.Now()
	expires := now.Add(time.Hour * 4).Unix()

	claims := jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": expires,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secret := os.Getenv("JWT_SECRET")
	// token must be a byte
	tokenStr, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Failed to sign token with secret")
	}
	return tokenStr
}

// Find the user through this.
func (h *AuthHandler) HandleAuthenticate(ctx context.Context) error {
	var params AuthParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := h.userStore.GetUserByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Fatal(err)
		}
		return err

	}
	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		log.Fatal(err)
	}

	// Creating our token from user
	resp := AuthResponse{
		User:  user,
		Token: CreateTokenFromUser(user),
	}
	return c.JSON(resp)

}

type AuthResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

type AuthParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
