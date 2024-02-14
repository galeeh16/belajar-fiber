package request

type CreateUserRequest struct {
	UserID          string `json:"user_id" validate:"required,min=3,max=50,unique_user_id"`
	Password        string `json:"password" validate:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Name            string `json:"name"  validate:"required,min=3,max=100"`
}
