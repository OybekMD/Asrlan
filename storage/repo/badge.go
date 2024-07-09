package repo

import "context"

type Badge struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	BadgeDate string `json:"badge_date"`
	BadgeType string `json:"badge_type"`
	Picture   string `json:"picture"`
}

type BadgeStorageI interface {
	Create(ctx context.Context, badge *Badge) (*Badge, error)
	Update(ctx context.Context, badge *Badge) (*Badge, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Badge, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*Badge, int64, error)
}
