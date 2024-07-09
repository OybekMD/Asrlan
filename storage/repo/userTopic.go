package repo

import "context"

type UserTopic struct {
	Id        string `json:"id"`
	Score     int64  `json:"score"`
	UserId    string `json:"user_id"`
	TopicId   int64  `json:"level_id"`
	CreatedAt string `json:"created_at"`
}

type UserTopicStorageI interface {
	Create(ctx context.Context, user_level *UserTopic) (bool, error)
}
