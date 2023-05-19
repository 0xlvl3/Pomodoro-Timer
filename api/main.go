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

	"github.com/0xlvl3/pomodoro-timer/api/db"
	"github.com/0xlvl3/pomodoro-timer/api/handles"
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
		todoStore = db.NewMongoTodoStore(client)

		// api
		app   = fiber.New()
		api   = app.Group("/api")
		apiv1 = app.Group("/api/v1", handles.JWTAuthentication(userStore))

		// handles
		userHandler = handles.NewUserHandler(userStore)
		todoHandler = handles.NewTodoHandler(todoStore)
		authHandler = handles.NewAuthHandler(userStore)
	)

	//TODO:  auth login and create
	api.Post("/auth", authHandler.HandleAuthenticate)

	//TODO: bring in handles
	api.Get("/test/:id", userHandler.HandleGetUserByID)
	api.Get("/user/:email", userHandler.HandleGetUserByEmail)
	api.Post("/user/create", userHandler.HandlePostUser)

	api.Post("/user/todo/add", todoHandler.HandleInsertTodo)
	apiv1.Get("/todo", todoHandler.HandleGetAllTodos)

	app.Listen(*lp)

}
