package request

type UserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required,min=3,max=20,alpha"`
	Age      int    `json:"age" validate:"required,min=1,max=120"`
}
