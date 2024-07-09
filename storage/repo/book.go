package repo

import "context"

type Book struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	BookFile  string `json:"book_file"`
	LevelId   int64  `json:"level_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type BookStorageI interface {
	Create(ctx context.Context, book *Book) (*Book, error)
	Update(ctx context.Context, book *Book) (*Book, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Book, error)
	GetAll(ctx context.Context, level_id string) ([]*Book, error)
}
