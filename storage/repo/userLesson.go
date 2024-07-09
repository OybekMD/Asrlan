package repo

import "context"

type UserLesson struct {
	Id        string `json:"id"`
	Score     int64  `json:"score"`
	UserId    string `json:"user_id"`
	LessonId  int64  `json:"level_id"`
	CreatedAt string `json:"created_at"`
}

type UserLessonStorageI interface {
	Create(ctx context.Context, userLesson *UserLesson) (bool, error)
}
