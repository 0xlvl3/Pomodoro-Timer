package db

import (
	"context"

	"github.com/0xlvl3/pomodoro-timer/api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoStore interface {
	InsertTodo(context.Context, *types.Todo) (*types.Todo, error)
	GetTodos(context.Context) ([]*types.Todo, error)
}

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

func (s *MongoTodoStore) InsertTodo(ctx context.Context, todo *types.Todo) (*types.Todo, error) {
	res, err := s.coll.InsertOne(ctx, todo)
	if err != nil {
		return nil, err
	}
	todo.ID = res.InsertedID.(primitive.ObjectID)

	return todo, nil
}

func (s *MongoTodoStore) GetTodos(ctx context.Context) ([]*types.Todo, error) {
	var todos []*types.Todo

	todo, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := todo.All(ctx, &todos); err != nil {
		return []*types.Todo{}, nil
	}

	return todos, nil
}
