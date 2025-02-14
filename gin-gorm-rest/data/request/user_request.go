package request

type CreateUserRequest struct {
	UserName string   `validate:"required,min=2,max=100" json:"username"`
	FullName string   `validate:"required,min=2,max=100" json:"fullname"`
	Email    string   `validate:"required,min=2,max=100" json:"email"`
	Password string   `validate:"required,min=2,max=100" json:"password"`
	Roles    []string `validate:"required,dive,min=2,max=50" json:"roles"`
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

type PermissionRoleRequest struct {
	Name  string   `json:"name" validate:"required,min=2,max=100"`
	Roles []string `json:"role" validate:"required,dive,min=2,max=50"`
}
