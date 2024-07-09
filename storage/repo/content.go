package repo

import "context"

type Content struct {
	Id            int64          `json:"id"`
	LessonId      int64          `json:"lesson_id"`
	Gentype       int64          `json:"gentype"`
	Title         string         `json:"title"`
	Question      string         `json:"question"`
	TextData      string         `json:"text_data"`
	ArrText       []string       `json:"arr_text"`
	CorrectAnswer int64          `json:"correct_answer"`
	Contentfiles  []*ContentFile `json:"contentfiles"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
}

type ContentStorageI interface {
	Create(ctx context.Context, badge *Content) (*Content, error)
	Update(ctx context.Context, badge *Content) (*Content, error)
	Delete(ctx context.Context, id string) (bool, error)
	Get(ctx context.Context, id string) (*Content, error)
	GetAll(ctx context.Context, id string) ([]*Content, int64, error)
}
