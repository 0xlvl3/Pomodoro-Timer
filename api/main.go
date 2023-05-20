package main

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

func main() {

	lp := flag.String("lp", ":8080", "listening port")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.URI))
	if err != nil {
		log.Fatal(err)
	}

	// unique usernames and emails
	db.SetupIndexes(client.Database(db.DBNAME).Collection("users"))

	var (
		// stores
		userStore = db.NewMongoUserStore(client)
		todoStore = db.NewMongoTodoStore(client)

		// api
		app  = fiber.New()
		api  = app.Group("/api")
		auth = app.Group("/auth", handles.JWTAuthentication(userStore))

		// handles
		userHandler = handles.NewUserHandler(userStore)
		todoHandler = handles.NewTodoHandler(todoStore)
		authHandler = handles.NewAuthHandler(userStore)
	)

	// login / create
	api.Post("/user/create", userHandler.HandlePostUser) //TODO: -- epass showing?
	api.Post("/login", authHandler.HandleAuthenticate)

	// user handles
	auth.Get("/user/:id", userHandler.HandleGetUserByID)
	auth.Get("/lookup/:email", userHandler.HandleGetUserByEmail)

	// todo handles
	auth.Post("/add/todo", todoHandler.HandleInsertTodo)
	auth.Get("/get/todo", todoHandler.HandleGetUserTodos)
	//auth.Get("/todo", todoHandler.HandleGetAllTodos) -- before we had specific user todos

	app.Listen(*lp)
}
