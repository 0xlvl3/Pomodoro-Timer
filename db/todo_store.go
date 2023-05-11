package db

import (
	"context"

	"github.com/0xlvl3/pomodoro-timer/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoStore interface {
	InsertTodo(context.Context, string, string) (*types.Todo, error)
}

// mongo
type MongoTodoStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoTodoStore(client *mongo.Client) *MongoTodoStore {
	return &MongoTodoStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("todos"),
	}
}

func (s *MongoTodoStore) InsertTodo(ctx context.Context, title string, description string) (*types.Todo, error) {
	todo := &types.Todo{
		Title:       title,
		Description: description,
	}

	res, err := s.coll.InsertOne(ctx, &todo)
	if err != nil {
		return nil, err
	}
	todo.ID = res.InsertedID.(primitive.ObjectID)

	return todo, nil
}