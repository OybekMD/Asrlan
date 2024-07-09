package models

import "asrlan-monolight/storage/repo"

type (
	// DashboardResponse struct {
	// 	Topics []*repo.Topic
	// }

	// TopicLessonResponse struct {
	// 	Id        int64  `json:"id"`
	// 	Name      string `json:"lesson_name"`
	// 	LevelId   int64  `json:"level_id"`
	// 	Topics []*repo.Topic
	// }

	DashboardResponse struct {
		Topics []Topics `json:"topics"`
	}
	LeaderboardResponse struct {
		Leaders []repo.Leaderboard `json:"leaders"`
	}

	Topics struct {
		TopicId    int64     `json:"topic_id"`
		TopicName  string    `json:"topic_name"`
		TopicScore int64     `json:"score"`
		Lessons    []Lessons `json:"lessons"`
	}
	Lessons struct {
		LessonId    int64  `json:"id"`
		LessonName  string `json:"lesson_name"`
		LessonScore int64  `json:"score"`
	}

	// DashboardListResponse struct {
	// 	Dashboards map[string][]*repo.Dashboard
	// }
)
