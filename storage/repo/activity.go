package repo

import "context"

type Activity struct {
	Id       int64  `json:"id"`
	Day      string `json:"day"`
	Score      int64 `json:"score"`
	LessonId int64  `json:"lesson_id"`
	UserId   string `json:"user_id"`
}

type ActivityStorageI interface {
	Create(ctx context.Context, badge *Activity) (*Activity, error)
	GetAllGroupedByMonth(ctx context.Context) (map[string][]*Activity, error)
	GetAllGroupedByChoice(ctx context.Context, choise string) (map[string][]*Activity, error)
}
