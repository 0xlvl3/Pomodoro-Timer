package main

import (
	"fmt"
	"log"

	"github.com/0xlvl3/pomodoro-timer/app"
)

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

func main() {
	loggedIn := false
	for !loggedIn {
		var input int
		fmt.Println("Welcome -- ")
		fmt.Println("Choose from the following options")
		fmt.Println("1 - Log in")
		fmt.Println("2 - Create account")
		fmt.Scanln(&input)

		if input == 1 {
			var email string
			var password string
			fmt.Println("-- Log in --")

			fmt.Printf("Email: ")
			//fmt.Scanln(&email)
			//email := ReadInput()
			fmt.Scanln(&email)

			fmt.Printf("Password: ")
			//password := ReadInput()
			fmt.Scanln(&password)

			username, err := app.Login(email, password)
			if username == "" {
				fmt.Println("Error with account log in")
				log.Fatal(err)
			} else {
				loggedIn = true
				app.Mainloop(username)
			}

		} else if input == 2 {
			var username string
			var email string
			var password string

			fmt.Printf("Choose a username: ")
			fmt.Scanln(&username)

			fmt.Printf("Choose a email: ")
			fmt.Scanln(&email)

			fmt.Printf("Choose a password: ")
			fmt.Scanln(&password)

			app.CreateUser(username, email, password)
			fmt.Println("Account created please go to log in!")
			continue
		} else {
			fmt.Println("Invalid option please enter 1 to log in or 2 to create an account")
			continue
		}
	}
}

//func StartMenu() {
//	//TODO: get this to work
//
//	fmt.Println("Create account? (c) - create user")
//	fmt.Println("Log in? (l) - log in")
//	fmt.Println("Or continue without one? (p) - proceed")
//	login := ReadUserInput(" ")
//	switch login {
//	case "c":
//		// login
//		username := ReadUserInput("username -- ")
//
//		email := ReadUserInput("email -- ")
//
//		password := ReadUserInput("password -- ")
//
//		user, err := userStore.NewUser(context.TODO(), username, email, password)
//		if err != nil {
//			log.Fatal(err)
//		}
//		fmt.Println(user)
//
//	case "l":
//		fmt.Println("\nLogin")
//	case "p":
//		fmt.Println("\nWelcome Guest")
//	}
//
//	StartPomodoroSession()
//
//}
