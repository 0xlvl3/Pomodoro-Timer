package db

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0xlvl3/pomodoro-timer/types"
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

// Menu for easy navigation
func (s *MongoPomodoroStore) NavigationMenu(nextAction, userInput string) {

	switch userInput {
	case "y":
		if nextAction == "study" {
			s.StartStudySession()
		} else {
			s.StartPomodoroBreak()
		}
	case "m":
		s.StartPomodoroSession()
	case "t":
		s.AddTodo()
	case "q":
		fmt.Println("Quitting...")
		os.Exit(2)
	}

	fmt.Println("\nwhere now?")
}

// StartPomodoroBreak a new break timer
func (s *MongoPomodoroStore) StartPomodoroBreak() {

	var duration int
	var input string
	fmt.Printf("How long do you want to break for: ")
	fmt.Scanf("%v", &duration)

	mins := duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo break starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo break starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: make quit function
	fmt.Println("q to quit at anytime")

	TimeLoop("Break", mins)

	fmt.Printf("\n\ngo to study (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &input)

	s.NavigationMenu("study", input)
}

// StartStudySession a new study timer
func (s *MongoPomodoroStore) StartStudySession() {

	var duration int
	var input string

	fmt.Printf("How long do you want to study for: ")
	fmt.Scanf("%v", &duration)

	mins := duration * 60

	// start loop over time stated
	if mins <= 60 {
		fmt.Printf("pomo starting for %d minute ..\n\n", mins/60)
	} else if mins > 60 {
		fmt.Printf("pomo starting for %d minutes ..\n\n", mins/60)
	}

	//TODO: add quit function
	fmt.Println("q to quit at anytime")

	TimeLoop("Study", mins)

	fmt.Printf("\n\ngo to break (y) yes, (m) menu or (q) quit: ")
	fmt.Scanf("%v", &input)

	s.NavigationMenu("break", input)
}

// Start is our init and welcome menu
func (s *MongoPomodoroStore) StartPomodoroSession() {
	//TODO: login user if wanted
	//TODO: username, password, stored in db
	input := " "
	for input != "q" {
		fmt.Printf("\nDo you want to start a pomo: (y) yes, (t) todo, (q) quit: ")

		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		s.NavigationMenu("study", input)
	}
}

// AddTodo will add a todo to a users db
func (s *MongoPomodoroStore) AddTodo() {
	todoStore := NewMongoTodoStore(s.client)

	fmt.Println("todo")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	//TODO: add time limit or num of pomos required

	todo, err := todoStore.InsertTodo(context.TODO(), title, description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("todo added :)", todo)

}

// TODO: list all todos
func (s *MongoPomodoroStore) ListTodos() *[]types.Todo {
	todoStore := NewMongoTodoStore(s.client)

	fmt.Println("todo list")

	var todos []types.Todo
	todo, err := todoStore.GetTodos(context.TODO())
	if err != nil {
		return nil, err
	}

	return todos, err

}
