package main

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

// clean up code

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/0xlvl3/pomodoro-timer/db"
	"github.com/0xlvl3/pomodoro-timer/handles"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// make db for users
// make db for todo

func main() {
	//
	lp := flag.String("lp", ":8080", "listening port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	//TODO: bring in stores
	var (
		// stores
		userStore = db.NewMongoUserStore(client)
		p         = db.NewMongoPomodoroStore(client)

		// api
		app    = fiber.New()
		listen = app.Group("/api")

		// handles
		userHandler = handles.NewUserHandler(userStore)
	)

	//TODO: bring in handles
	listen.Get("/user", userHandler.HandleGetUserByEmail)

	//TODO: make it work

	//TODO: Remove this into handles
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

	app.Listen(*lp)

}
