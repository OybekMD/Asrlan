package models

import "asrlan-monolight/storage/repo"

type (
	ContentCreate struct {
		LessonId      int64    `json:"lesson_id"`
		Gentype       int64    `json:"gentype"`
		Title         string   `json:"title"`
		Question      string   `json:"question"`
		TextData      string   `json:"text_data"`
		ArrText       []string `json:"arr_text"`
		CorrectAnswer int64    `json:"correct_answer"`
	}

	ContentRequest struct {
		Id string `json:"id"`
	}

	ContentResponse struct {
		Id            int64               `json:"id"`
		LessonId      int64               `json:"lesson_id"`
		Gentype       int64               `json:"gentype"`
		Title         string              `json:"title"`
		Question      string              `json:"question"`
		TextData      string              `json:"text_data"`
		ArrText       []string            `json:"arr_text"`
		CorrectAnswer int64               `json:"correct_answer"`
		Contentfiles  []*repo.ContentFile `json:"contentfiles"`
		CreatedAt     string              `json:"created_at"`
		UpdatedAt     string              `json:"updated_at"`
	}

	ContentListResponse struct {
		Contents []*repo.Content
		Count    int64 `json:"count"`
	}

	ContentUpdate struct {
		Id            int64               `json:"id"`
		LessonId      int64               `json:"lesson_id"`
		Gentype       int64               `json:"gentype"`
		Title         string              `json:"title"`
		Question      string              `json:"question"`
		TextData      string              `json:"text_data"`
		ArrText       []string            `json:"arr_text"`
		CorrectAnswer int64               `json:"correct_answer"`
	}
)
