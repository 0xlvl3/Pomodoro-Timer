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
	user := c.Locals("user").(*types.User)
	fmt.Println("iiiiii---", user.ID)
	//TODO: add time limit or num of pomos required

	var todo *types.Todo
	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	todo.UserID = user.ID

	addedTodo, err := h.todoStore.InsertTodo(c.Context(), todo)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("- Todo added - \nTitle: %v - \nDescription: %v\n userID - %v\n", addedTodo.Title, addedTodo.Description, addedTodo.UserID)
	return c.JSON(addedTodo)
}

//func (h *TodoHandler) HandleGetAllTodos(c *fiber.Ctx) error {
//
//	todos, err := h.todoStore.GetAllTodos(c.Context())
//	if err != nil {
//		return err
//	}
//
//	fmt.Println(todos)
//
//	return c.JSON(todos)
//}

func (h *TodoHandler) HandleGetUserTodos(c *fiber.Ctx) error {
	user := c.Locals("user").(*types.User)
	userID := user.ID

	todos, err := h.todoStore.GetTodosByUserID(c.Context(), userID)
	if err != nil {
		return err
	}

	fmt.Println(todos)
	return c.JSON(todos)
}
