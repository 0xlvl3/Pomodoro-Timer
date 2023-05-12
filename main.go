package main

// ideas
// loop handled by own function

// add todo to aspect where user can add a list of todo
// todo then can be studied through pomo

// mongodb - user accs, saves data
// don't require mongodb can be optional

// understand a bit more on styling through lipgloss
// - figure way that it won't reprint on the same line with \r "lipgloss"

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

// clean up code

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/0xlvl3/pomodoro-timer/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// make db for users
// make db for todo

func main() {

	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("client starting %v \n\n", client)

	userStore := db.NewMongoUserStore(client)
	p := db.NewMongoPomodoroStore(client)

	fmt.Println("Create account? (c) - create user")
	fmt.Println("Log in? (l) - log in")
	fmt.Println("Or continue without one? (p) - proceed")
	login := db.ReadUserInput(" ")
	switch login {
	case "c":
		// login
		username := db.ReadUserInput("username -- ")

		email := db.ReadUserInput("email -- ")

		password := db.ReadUserInput("password -- ")

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

	p.StartPomodoroSession()

}
