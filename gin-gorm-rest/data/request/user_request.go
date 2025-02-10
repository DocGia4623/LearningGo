package request

type CreateUserRequest struct {
	UserName string `validate:"required,min=2,max=100" json:"username"`
	FullName string `validate:"required,min=2,max=100" json:"fullname"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	Role     string `validate:"required,min=2,max=100" json:"role"`
}
type UpdateUserRequest struct {
	Id       uint   `validate:"required"`
	UserName string `validate:"required,min=2,max=100" json:"username"`
	FullName string `validate:"required,min=2,max=100" json:"fullname"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
	Role     string `validate:"required,min=2,max=100" json:"role"`
}

type LoginRequest struct {
	UserName string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
