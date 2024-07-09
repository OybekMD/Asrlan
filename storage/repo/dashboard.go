package repo

import "context"

type Navbar struct {
	Message         []string   `json:"messages"`
	Streak          int64      `json:"streak"`
	Score           int64      `json:"score"`
	ActiveLanguages []*Language `json:"active_languages"`
}

type Dashboard struct {
	Topics []Topics `json:"topics"`
}

type Topics struct {
	TopicId    int64     `json:"topic_id"`
	TopicName  string    `json:"topic_name"`
	TopicScore int64     `json:"score"`
	Lessons    []Lessons `json:"lessons"`
}
type Lessons struct {
	LessonId    int64  `json:"id"`
	LessonType  string `json:"lesson_type"`
	LessonScore int64  `json:"score"`
}

type Leaderboard struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Score    int64  `json:"score"`
}

type DashboardStorageI interface {
	GetNavbar(ctx context.Context, userId string) (*Navbar, error)
	GetDashboard(ctx context.Context, userId string, language_id, level_id int64) (*Dashboard, error)
	GetLeaderboard(ctx context.Context, period, level_id string) ([]*Leaderboard, error)
	UpCreateLevel(ctx context.Context, userLevel *UserLevel) (bool, error)
	UpCreateTopic(ctx context.Context, userTopic *UserTopic) (bool, error)
}
