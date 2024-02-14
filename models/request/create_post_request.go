package request

type CreatePostRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=3,max=1000"`
}
