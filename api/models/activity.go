package models

import "asrlan-monolight/storage/repo"

type (
	ActivityCreate struct {
		Day      string `json:"day"`
		Score    int64  `json:"score"`
		LessonId int64  `json:"lesson_id"`
		UserId   string `json:"user_id"`
	}

	ActivityRequest struct {
		Id string `json:"id"`
	}

	ActivityResponse struct {
		Id       int64  `json:"id"`
		Day      string `json:"day"`
		Score    int64  `json:"score"` //It for user_lesson
		LessonId int64  `json:"lesson_id"`
		UserId   string `json:"user_id"`
	}

	ActivityListResponse struct {
		Activitys map[string][]*repo.Activity
	}
)
