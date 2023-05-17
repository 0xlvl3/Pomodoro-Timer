package main

// bring in tea model
// has selections for pre-defined study times 15, 25, 40, 60
// has slected for pre-defined break times 5, 10, 15
// todos have own selection board

// clean up code

import (
	"context"
	"flag"
	"log"

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
		userStore     = db.NewMongoUserStore(client)
		pomodoroStore = db.NewMongoPomodoroStore(client)

		// api
		app    = fiber.New()
		listen = app.Group("/api")

		// handles
		userHandler     = handles.NewUserHandler(userStore)
		pomodoroHandler = handles.NewPomodoroHandler(pomodoroStore)
	)

	//TODO: bring in handles
	listen.Get("/user", userHandler.HandleGetUserByEmail)
	listen.Get("/pomo", pomodoroHandler.NavigationMenu("test", "test"))

}
