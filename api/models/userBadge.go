package models

import "asrlan-monolight/storage/repo"

type (
	UserBadgeCreate struct {
		UserId  string `json:"user_id"`
		BadgeId string `json:"badge_id"`
	}

	UserBadgeDelete struct {
		UserId  string `json:"user_id"`
		BadgeId string `json:"badge_id"`
	}

	UserBadgeRequest struct {
		UserId string `json:"user_id"`
	}

	UserBadgeListResponse struct {
		UserBadges []*repo.Badge
		Count      int64 `json:"count"`
	}
)
