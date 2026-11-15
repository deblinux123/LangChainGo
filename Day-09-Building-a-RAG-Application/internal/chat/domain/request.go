package domain

type ChatRequest struct {
	UserID  uint   `json:"user_id"`
	Message string `json:"message"`
}
