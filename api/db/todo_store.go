package db

import (
	"context"
	"log"

	"github.com/0xlvl3/pomodoro-timer/api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TodoStore interface {
	InsertTodo(context.Context, *types.Todo) (*types.Todo, error)
	//	GetAllTodos(context.Context) ([]*types.Todo, error)
	GetTodosByUserID(context.Context, primitive.ObjectID) ([]*types.Todo, error)
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

//func (s *MongoTodoStore) GetAllTodos(ctx context.Context) ([]*types.Todo, error) {
//
//	// bson.M{} == make(map[string]any)
//	cur, err := s.coll.Find(ctx, make(map[string]any))
//	if err != nil {
//		return nil, err
//	}
//
//	var todos []*types.Todo
//	if err := cur.All(ctx, &todos); err != nil {
//		return nil, err
//	}
//
//	return todos, nil
//}

func (s *MongoTodoStore) GetTodosByUserID(ctx context.Context, userID primitive.ObjectID) ([]*types.Todo, error) {

	filter := bson.M{"userID": userID}
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var todos []*types.Todo
	if err := cur.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// used to create unique usernames and emails
func SetupIndexes(collection *mongo.Collection) {
	usernameIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	emailIndexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateMany(context.Background(), []mongo.IndexModel{usernameIndexModel, emailIndexModel})
	if err != nil {
		log.Fatal(err)
	}
}
