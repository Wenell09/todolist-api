package todo

type ReqTodo struct {
	TodoTitle   string `json:"todo_title"`
	TodoDesc    string `json:"todo_desc"`
	UserId      string `json:"user_id"`
	PrioritasId string `json:"prioritas_id"`
	StatusId    string `json:"status_id"`
}
