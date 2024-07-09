package repo

import "context"

type Lesson struct {
	Id         int64  `json:"id"`
	LessonType string `json:"lesson_type"`
	TopicId    int64  `json:"topic_id"`
	TopicName  string `json:"topic_name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type LessonStorageI interface {
	Create(ctx context.Context, badge *Lesson) (*Lesson, error)
	Update(ctx context.Context, badge *Lesson) (*Lesson, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Lesson, error)
	GetAll(ctx context.Context, lesson_id string) ([]*Lesson, error)
}
