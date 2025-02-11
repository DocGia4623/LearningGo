package response

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type DeviceResponse struct {
	ID       uint         `json:"id"`
	Name     string       `json:"fullname"`
	Price    int          `json:"price"`
	Quantity int          `json:"quantity"`
	User     UserResponse `json:"user"`
}
