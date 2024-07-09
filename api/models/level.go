package models

import "asrlan-monolight/storage/repo"

type (
	LevelCreate struct {
		Name       string `json:"name"`
		LanguageId int64  `json:"language_id"`
	}

	LevelRequest struct {
		Id string `json:"id"`
	}

	LevelResponse struct {
		Id           int64  `json:"id"`
		Name         string `json:"name"`
		LanguageId   int64  `json:"language_id"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	LevelListResponse struct {
		Levels []*repo.Level
		Count  int64 `json:"count"`
	}

	LevelForRegisterResponse struct {
		Levels []*repo.LevelForRegister
	}

	LevelForCourseResponse struct {
		Levels []*repo.LevelForCourse
	}

	LevelUpdate struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}
)
