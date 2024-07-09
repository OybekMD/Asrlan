package repo

import (
	"context"
)

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Bio          string `json:"bio"`
	BirthDay     string `json:"birth_day"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Avatar       string `json:"avatar"`
	Coint        int64  `json:"coint"`
	Score        int64  `json:"score"`
	LanguageId   int64  `json:"language_id"`
	LevelId      int64  `json:"level_id"`
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type LoginPassword struct {
	UserId   string `json:"user_id"`
	Role     string `json:"role"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResetPassword struct {
	Otp         string `json:"otp"`
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

type LoginStorageI interface {
	Login(ctx context.Context, login string) (*LoginResponse, error)
	SavePassword(ctx context.Context, req *LoginPassword) (*LoginPassword, error)
	ResetPassword(ctx context.Context, req *ResetPassword) (*LoginResponse, error)
	GetUserByLogin(ctx context.Context, login string) (id string, role string, err error)
	SaveRefresh(ctx context.Context, role, id, refresh string) (bool, error)
}
