package models

import "asrlan-monolight/storage/repo"

type (
	LanguageCreate struct {
		Name    string `json:"name"`
		Picture string `json:"picture"`
	}

	LanguageRequest struct {
		Id string `json:"id"`
	}

	LanguageResponse struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		Picture   string `json:"picture"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	LanguageListResponse struct {
		Languages []*repo.Language
		Count     int64 `json:"count"`
	}

	LanguageForRegisterResponse struct {
		Languages []*repo.RegisterLanguage
	}

	LanguageUpdate struct {
		Id        int64  `json:"id"`
		Name      string `json:"name"`
		Picture   string `json:"picture"`
	}
)
