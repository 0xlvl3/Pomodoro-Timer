package handles

import (
	"fmt"

	"github.com/0xlvl3/pomodoro-timer/api/db"
	"github.com/0xlvl3/pomodoro-timer/api/types"
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

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	//TODO: fix this error
	//	if errors := params.Validate(); len(errors) > 0 {
	//		return c.JSON(errors)
	//	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
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
