package todo

import (
	"time"
)

type Todo struct {
	TodoId      string
	TodoTitle   string
	TodoDesc    string
	UserId      string
	PrioritasId string
	CreatedAt   time.Time
}
