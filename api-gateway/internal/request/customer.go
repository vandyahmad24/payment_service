package request

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"  binding:"required,min=6"`
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"  binding:"required,min=6"`
	Name     string `json:"name" binding:"required"`
}
