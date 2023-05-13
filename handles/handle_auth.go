package handles

import "github.com/0xlvl3/pomodoro-timer/db"

type AuthHandler struct {
	userStore db.UserStore
}
