package todo

import (
	"time"
)

type Todo struct {
	TodoId      string    `json:"todo_id"`
	TodoTitle   string    `json:"todo_title"`
	TodoDesc    string    `json:"todo_desc"`
	UserId      string    `json:"user_id"`
	PrioritasId string    `json:"prioritas_id"`
	StatusId    string    `json:"status_id"`
	CreatedAt   time.Time `json:"created_at"`
}
