package user

type ReqUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
