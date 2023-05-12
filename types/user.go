package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost     = 12
	minUsernameLen = 2
	minPasswordLen = 10
)

type User struct {
	ID                primitive.ObjectID
	Username          string
	Email             string
	EncryptedPassword string
}

type CreateUserParams struct {
	Username string
	Email    string
	Password string
}

func NewUserFromParams(params CreateUserParams) (*User, error) {

	// encrypt our password
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	// create user through params
	return &User{
		Username:          params.Username,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil

}

func IsValidPassword(encpw, pw string) bool {

	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func (params CreateUserParams) Validate() map[string]string {
	errors := map[string]string{}
	if len(params.Username) < minUsernameLen {
		errors["firstName"] = fmt.Sprintf("firstName length should be at least %d characters", minUsernameLen)
	}
	if len(params.Password) < minPasswordLen {
		errors["password"] = fmt.Sprintf("password length should be at least %d characters", minPasswordLen)
	}
	if !isEmailValid(params.Email) {
		errors["email"] = fmt.Sprintf("%s email is invalid", params.Email)
	}
	return errors
}
