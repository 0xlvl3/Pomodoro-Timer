package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type PomodoroStore interface {
	StartPomodoroSession()
	StartStudySession()
	StartPomodoroBreak()
	NavigationMenu(string, string)
}

type MongoPomodoroStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoPomodoroStore(client *mongo.Client) *MongoPomodoroStore {
	return &MongoPomodoroStore{
		client: client,
		coll:   client.Database(DBNAME).Collection("pomodoro"),
	}

}
