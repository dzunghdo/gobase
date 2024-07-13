package dto

type LoginRequest struct {
	Username string `json:"username",required`
	Password string `json:"password",required`
}

type RegisterRequest struct {
	Username string `json:"username",required`
	Password string `json:"password",required`
	Email    string `json:"email",required`
}
