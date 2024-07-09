package repo

import "context"

type User struct {
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
	RefreshToken string `json:"refresh_token"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type UserUpdatePassword struct {
	Id          string `json:"id"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type UserStorageI interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*User, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*User, error)
	CheckField(ctx context.Context, field, value string) (bool, error)
	CheckUsername(ctx context.Context, id, username string) (bool, error)
	UpdatePassword(ctx context.Context, newUser *UserUpdatePassword) (bool, error)
}
