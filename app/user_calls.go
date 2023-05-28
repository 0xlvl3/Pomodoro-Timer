package app

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/0xlvl3/pomodoro-timer/api/types"
)

// will contain all user api call methods

func GetUserByID() {

	resp, err := http.Get("http://localhost:8080/api/test/645df0bd004e4e0b5064590f")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

	var user types.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Email, user.Username, user.ID)

}

func GetUserByEmail() {
	var email string
	fmt.Printf("What is your email: ")
	fmt.Scanln(&email)

	resp, err := http.Get("http://localhost:8080/api/user/" + email)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user types.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(user.Email, user.Username, user.ID)

}

func CreateUser(username, email, password string) {
	user := &types.CreateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8080/api/user/create", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Trim off the newline at the end
	return input
}
