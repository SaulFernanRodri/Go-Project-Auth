package models

type AuthUser struct {
	ID       uint   `gorm:"primaryKey;autoIncriment" json:"id"`
	Username string `gorm:"unique" json:"username" binding:"required,min=4"`
	Email    string `gorm:"unique" json:"email" binding:"required,email"`
	Password string `gorm:"size:100" json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
