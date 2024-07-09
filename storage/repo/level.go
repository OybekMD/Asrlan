package repo

import "context"

type Level struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	RealLevel    int64  `json:"real_level"`
	Picture      string `json:"picture"`
	LanguageId   int64  `json:"language_id"`
	LanguageName string `json:"language_name"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type LevelForRegister struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	RealLevel int64  `json:"real_level"`
	Picture   string `json:"picture"`
}

type LevelForCourse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Score     int64  `json:"score"`
	RealLevel int64  `json:"real_level"`
	Picture   string `json:"picture"`
}

type LevelStorageI interface {
	Create(ctx context.Context, badge *Level) (*Level, error)
	Update(ctx context.Context, badge *Level) (*Level, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Level, error)
	GetAll(ctx context.Context, page, limit uint64) ([]*Level, int64, error)
	GetAllForRegister(ctx context.Context, language_id string) ([]*LevelForRegister, error)
	GetAllForCourses(ctx context.Context, user_id, language_id string) ([]*LevelForCourse, error)
}
