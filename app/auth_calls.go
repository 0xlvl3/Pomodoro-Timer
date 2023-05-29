package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/0xlvl3/pomodoro-timer/api/handles"
	"github.com/0xlvl3/pomodoro-timer/api/types"
)

type LoginResponse struct {
	User  types.User `json:"user"`
	Token string     `json:"token"`
}

func Login(email, password string) (string, error) {

	loginUser := handles.AuthParams{
		Email:    email,
		Password: password,
	}

	jsonUser, err := json.Marshal(loginUser)
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	fmt.Println(loginResp.Token)
	fmt.Println(loginResp.User.Username)

	return loginResp.User.Username, nil
}
