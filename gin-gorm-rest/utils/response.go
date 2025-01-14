package utils

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type DeviceResponse struct {
	ID       uint         `json:"id"`
	Name     string       `json:"name"`
	Price    int          `json:"price"`
	Quantity int          `json:"quantity"`
	User     UserResponse `json:"user"`
}
