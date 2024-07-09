package models

import "asrlan-monolight/storage/repo"

type (
	TopicCreate struct {
		Name string `json:"topic_name"`
	}

	TopicRequest struct {
		Id string `json:"id"`
	}

	TopicResponse struct {
		Id        int64  `json:"id"`
		Name      string `json:"topic_name"`
		LevelId   int64  `json:"level_id"`
		LevelName string `json:"level_name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	TopicListResponse struct {
		Topics []*repo.Topic
		Count  int64 `json:"count"`
	}

	TopicUpdate struct {
		Id   int64  `json:"id"`
		Name string `json:"topic_name"`
	}
)
