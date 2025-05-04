package todo

import (
	"time"
)

type ResTodo struct {
	TodoId        string    `json:"todo_id"`
	TodoTitle     string    `json:"todo_title"`
	TodoDesc      string    `json:"todo_desc"`
	UserId        string    `json:"user_id"`
	Username      string    `json:"username"`
	PrioritasId   string    `json:"prioritas_id"`
	PrioritasName string    `json:"prioritas_name"`
	CreatedAt     time.Time `json:"created_at"`
}
