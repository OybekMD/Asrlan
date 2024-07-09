package repo

import "context"

type Admin struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type AdminStorageI interface {
	Create(ctx context.Context, email string) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, email string) (bool, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*Admin, error)
}
