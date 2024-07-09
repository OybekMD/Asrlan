package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
)

type userLessonRepo struct {
	db *sqlx.DB
}

func NewUserLesson(db *sqlx.DB) repo.UserLessonStorageI {
	return &userLessonRepo{
		db: db,
	}
}

// This function create userLesson in postgres
func (s *userLessonRepo) Create(ctx context.Context, userLesson *repo.UserLesson) (bool, error) {
	pp.Println(userLesson)
	queryExist := `
	UPDATE
		user_lesson
	SET
		score = $1
	WHERE
		user_id = $2
		AND lesson_id = $3
	`
	result, err := s.db.ExecContext(ctx,
		queryExist,
		userLesson.Score,
		userLesson.UserId,
		userLesson.LessonId,
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
	INSERT INTO user_lesson (
		score,
		user_id,
		lesson_id
	)
	VALUES ($1, $2, $3)
	`

	_, err = s.db.ExecContext(
		ctx,
		query,
		userLesson.Score,
		userLesson.UserId,
		userLesson.LessonId,
	)

	if err != nil {
		log.Println("Eror creating userTopic in postgres method", err.Error())
		return false, err
	}

	return true, err
}
