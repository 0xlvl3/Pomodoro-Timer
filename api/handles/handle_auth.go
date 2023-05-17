package handles

import (
	"context"
	"fmt"
	"log"

	"github.com/0xlvl3/pomodoro-timer/api/db"
)

type AuthHandler struct {
	userStore db.UserStore
}

func StartMenu() {
	//TODO: get this to work

	fmt.Println("Create account? (c) - create user")
	fmt.Println("Log in? (l) - log in")
	fmt.Println("Or continue without one? (p) - proceed")
	login := ReadUserInput(" ")
	switch login {
	case "c":
		// login
		username := ReadUserInput("username -- ")

		email := ReadUserInput("email -- ")

		password := ReadUserInput("password -- ")

		user, err := userStore.NewUser(context.TODO(), username, email, password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)

	case "l":
		fmt.Println("\nLogin")
	case "p":
		fmt.Println("\nWelcome Guest")
	}

	StartPomodoroSession()

}
