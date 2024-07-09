package repo

import "context"

type Language struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegisterLanguage struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	UserCount int64  `json:"user_count"`
}

type LanguageStorageI interface {
	Create(ctx context.Context, badge *Language) (*Language, error)
	Update(ctx context.Context, badge *Language) (*Language, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Language, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*Language, int64, error)
	GetAllForRegister(ctx context.Context) ([]*RegisterLanguage, error)
}
