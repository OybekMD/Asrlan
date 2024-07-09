package repo

import "context"

type UserLevel struct {
	Id        string `json:"id"`
	Score     int64  `json:"score"`
	Status    bool   `json:"status"`
	UserId    string `json:"user_id"`
	LevelId   int64  `json:"level_id"`
	CreatedAt string `json:"created_at"`
}

type UserLevelStorageI interface {
	Create(ctx context.Context, user_level *UserLevel) (bool, error)
	GetActive(ctx context.Context, userId string) (*UserLevel, error)
	GetAll(ctx context.Context, userId string) ([]*UserLevel, error)
}
