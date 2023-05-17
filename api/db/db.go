package db

const (
	URI    = "mongodb+srv://admin:admin@pomo.xnefbgd.mongodb.net/"
	DBNAME = "pomo"
)

type Store struct {
	User UserStore
	Todo TodoStore
}
