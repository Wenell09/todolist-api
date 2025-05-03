package user

import "time"

type User struct {
	UserId    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt time.Time
}
