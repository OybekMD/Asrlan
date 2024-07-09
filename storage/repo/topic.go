package repo

import "context"

type Topic struct {
	Id        int64  `json:"id"`
	Name string `json:"topic_name"`
	LevelId   int64  `json:"level_id"`
	LevelName string `json:"level_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TopicStorageI interface {
	Create(ctx context.Context, badge *Topic) (*Topic, error)
	Update(ctx context.Context, badge *Topic) (*Topic, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Topic, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*Topic, int64, error)
}
