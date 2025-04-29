package request

type UserRequest struct {
	Password string `json:"password" binding:"required,min=6,max=20,containsany=@!#$%&*"`
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required,min=4,max=20,alpha"`
	Age      int    `json:"age" binding:"required,min=1,max=100"`
}
