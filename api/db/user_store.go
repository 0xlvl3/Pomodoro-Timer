package db

import (
	"context"
	"fmt"
	"log"

	"github.com/0xlvl3/pomodoro-timer/api/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	NewUser(context.Context, string, string, string) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("users"),
	}
}

func (s *MongoUserStore) NewUser(ctx context.Context, username, email, password string) (*types.User, error) {
	var params types.CreateUserParams
	params.Email = email
	params.Username = username
	params.Password = password

	if errors := params.Validate(); len(errors) > 0 {
		log.Fatal(errors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return nil, err
	}
	insertUser, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	fmt.Println("user added -- ", insertUser)

	return nil, nil
}

func (s *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user *types.User
	if err := s.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}
