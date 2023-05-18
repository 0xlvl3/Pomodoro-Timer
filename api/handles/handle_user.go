package handles

import (
	"fmt"

	"github.com/0xlvl3/pomodoro-timer/api/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUserByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := h.userStore.GetUserByEmail(c.Context(), email)
	if err != nil {
		return err
	}

	fmt.Printf("%+v", user)
	return c.JSON(user)

}
