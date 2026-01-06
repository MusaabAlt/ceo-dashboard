package dto

import "github.com/google/uuid"

type RegisterRequest struct {
	CompanyName string `json:"company_name"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token   string      `json:"token"`
	User    UserInfo    `json:"user"`
	Company CompanyInfo `json:"company"`
}

type UserInfo struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	IsActive bool      `json:"is_active"`
}

type CompanyInfo struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	DefaultCurrency string    `json:"default_currency"`
	Timezone        string    `json:"timezone"`
}
