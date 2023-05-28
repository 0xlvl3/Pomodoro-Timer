package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/0xlvl3/pomodoro-timer/api/handles"
)

type LoginResponse struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

func Login(email, password string) {

	loginUser := handles.AuthParams{
		Email:    email,
		Password: password,
	}

	jsonUser, err := json.Marshal(loginUser)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/login", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var loginResp LoginResponse
	err = json.Unmarshal(body, &loginResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	fmt.Println(loginResp.Token)

}
