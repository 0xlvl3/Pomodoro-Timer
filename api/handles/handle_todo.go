package handles

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/0xlvl3/pomodoro-timer/api/db"
	"github.com/0xlvl3/pomodoro-timer/api/types"
)

type TodoHandler struct {
	todoStore db.TodoStore
}

func NewTodoHandler(todoStore *db.TodoStore) *TodoHandler {
	return &TodoHandler{
		todoStore: *todoStore,
	}
}

// AddTodo will add a todo to a users db
func (h *TodoHandler) AddTodo() {

	fmt.Println("todo")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	//TODO: add time limit or num of pomos required

	todo, err := h.todoStore.InsertTodo(context.TODO(), title, description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("todo added :)", todo)

}

func (h *TodoHandler) ListTodos() ([]*types.Todo, error) {

	fmt.Println("todo list")

	todos, err := h.todoStore.GetTodos(context.TODO())
	if err != nil {
		return nil, err
	}

	for i, todo := range todos {

		fmt.Printf("%d. Todo \n--- Title: %v \n--- Description: %v \n\n", i, todo.Title, todo.Description)
	}

	return todos, nil

}
