package dto

type CreateTaskRequest struct {
	UserID uint   `json:"user_id"`
	Title  string `json:"title"`
}

type UpdateTaskRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
