package models

import "asrlan-monolight/storage/repo"

type (
	BadgeCreate struct {
		Name      string `json:"name"`
		BadgeDate string `json:"badge_date"`
		BadgeType string `json:"badge_type"`
		Picture   string `json:"picture"`
	}

	BadgeRequest struct {
		Id string `json:"id"`
	}

	BadgeResponse struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		BadgeDate string `json:"badge_date"`
		BadgeType string `json:"badge_type"`
		Picture   string `json:"picture"`
	}

	BadgeListResponse struct {
		Badges []*repo.Badge
		Count  int64 `json:"count"`
	}

	BadgeUpdate struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		BadgeDate string `json:"badge_date"`
		BadgeType string `json:"badge_type"`
		Picture   string `json:"picture"`
	}
)
