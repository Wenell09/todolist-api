package todo

import (
	"time"

	"github.com/Wenell09/todolist-api/models/prioritas"
)

type Todo struct {
	TodoId      string              `json:"todo_id"`
	TodoTitle   string              `json:"todo_title"`
	TodoDesc    string              `json:"todo_desc"`
	UserId      string              `json:"user_id"`
	PrioritasId string              `json:"-"`
	Prioritas   prioritas.Prioritas `json:"prioritas" gorm:"foreignKey:PrioritasId"`
	CreatedAt   time.Time           `json:"created_at"`
}
