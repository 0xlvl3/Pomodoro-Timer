package handles

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0xlvl3/pomodoro-timer/types"
)

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

func (s *MongoPomodoroStore) ListTodos() ([]*types.Todo, error) {
	todoStore := NewMongoTodoStore(s.client)

	fmt.Println("todo list")

	todos, err := todoStore.GetTodos(context.TODO())
	if err != nil {
		return nil, err
	}

	for i, todo := range todos {

		fmt.Printf("%d. Todo \n--- Title: %v \n--- Description: %v \n\n", i, todo.Title, todo.Description)
	}

	return todos, nil

}
