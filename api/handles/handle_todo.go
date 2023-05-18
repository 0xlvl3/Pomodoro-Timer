package handles

import (
	"fmt"
	"log"

	"github.com/0xlvl3/pomodoro-timer/api/db"
	"github.com/0xlvl3/pomodoro-timer/api/types"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoStore db.TodoStore
}

func NewTodoHandler(todoStore db.TodoStore) *TodoHandler {
	return &TodoHandler{
		todoStore: todoStore,
	}
}

// AddTodo will add a todo to a users db
func (h *TodoHandler) HandleInsertTodo(c *fiber.Ctx) error {
	//TODO: add time limit or num of pomos required

	var todo *types.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	addedTodo, err := h.todoStore.InsertTodo(c.Context(), todo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("- Todo added - \nTitle: %v - \nDescription: %v\n", addedTodo.Title, addedTodo.Description)
	return c.JSON(addedTodo)
}

func (h *TodoHandler) ListTodos(c *fiber.Ctx) ([]*types.Todo, error) {

	fmt.Println("todo list")

	todos, err := h.todoStore.GetTodos(c.Context())
	if err != nil {
		return nil, err
	}

	for i, todo := range todos {

		fmt.Printf("%d. Todo \n--- Title: %v \n--- Description: %v \n\n", i, todo.Title, todo.Description)
	}

	return todos, nil

}
