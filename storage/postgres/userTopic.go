package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type userTopicRepo struct {
	db *sqlx.DB
}

func NewUserTopic(db *sqlx.DB) repo.UserTopicStorageI {
	return &userTopicRepo{
		db: db,
	}
}

func (s *userTopicRepo) Create(ctx context.Context, userTopic *repo.UserTopic) (bool, error) {
	queryExist := `
	UPDATE
		user_lesson
	SET
		score = $1
	WHERE
		user_id = $2
		AND topic_id = $3
	`
	result, err := s.db.ExecContext(ctx,
		queryExist,
		userTopic.Score,
		userTopic.UserId,
		userTopic.TopicId,
	)
	if err != nil {
		log.Println("Error deleting user", err.Error())
		return false, err
	}

	rowEffect, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect != 0 {
		return true, nil
	}

	query := `
	INSERT INTO user_topic (
		score,
		user_id,
		topic_id
	)
	VALUES ($1, $2, $3)
	`

	_, err = s.db.ExecContext(
		ctx,
		query,
		userTopic.Score,
		userTopic.UserId,
		userTopic.TopicId,
	)

	if err != nil {
		log.Println("Eror creating userTopic in postgres method", err.Error())
		return false, err
	}

	return true, err
}