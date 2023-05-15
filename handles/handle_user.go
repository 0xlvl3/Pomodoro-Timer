package handles

import (
	"github.com/0xlvl3/pomodoro-timer/db"
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

	return c.JSON(user)

}
