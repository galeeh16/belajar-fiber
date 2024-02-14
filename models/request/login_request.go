package request

type LoginRequest struct {
	UserID string `json:"user_id" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}