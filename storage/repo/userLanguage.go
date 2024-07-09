package repo

import "context"

type UserLanguage struct {
	Id         string `json:"id"`
	Score      int64  `json:"score"`
	Status     bool   `json:"status"`
	UserId     string `json:"user_id"`
	LanguageId int64  `json:"language_id"`
	CreatedAt  string `json:"created_at"`
}

type UserLanguageStorageI interface {
	Create(ctx context.Context, user_language *UserLanguage) (bool, error)
	GetActive(ctx context.Context, userId string) (*UserLanguage, error)
	GetAll(ctx context.Context, userId string) ([]*UserLanguage, error)
}
