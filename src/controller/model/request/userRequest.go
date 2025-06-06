package request

type UserRequest struct {
	Password string `json:"password" binding:"required,min=6,max=20,containsany=@!#$%&*"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=4,max=20"`
	Age      int8   `json:"age" binding:"required,min=1,max=100"`
}

type UserUpdateRequest struct {
	Name string `json:"name" binding:"omitempty,min=4,max=100" example:"John Doe"`
	Age  int8   `json:"age" binding:"omitempty,min=1,max=140" example:"30"`
}
