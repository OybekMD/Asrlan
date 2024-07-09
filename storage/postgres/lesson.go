package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type lessonRepo struct {
	db *sqlx.DB
}

func NewLesson(db *sqlx.DB) repo.LessonStorageI {
	return &lessonRepo{
		db: db,
	}
}

// This function create lesson in postgres
func (s *lessonRepo) Create(ctx context.Context, lesson *repo.Lesson) (*repo.Lesson, error) {
	query := `
	INSERT INTO lessons(
		lesson_type,
		topic_id
	)
	VALUES ($1, $2) 
	RETURNING 
		id,
		created_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		lesson.LessonType,
		lesson.TopicId).Scan(&lesson.Id, &lesson.CreatedAt)
	if err != nil {
		log.Println("Eror creating lesson in postgres method", err.Error())
		return nil, err
	}

	return lesson, nil
}

// This function update lesson info from postgres
func (s *lessonRepo) Update(ctx context.Context, newLesson *repo.Lesson) (*repo.Lesson, error) {
	query := `
	UPDATE
		lessons
	SET
		lesson_type=$1,
		updated_at=CURRENT_TIMESTAMP
	WHERE
		id=$2
	AND deleted_at IS NULL
	RETURNING
		created_at,
		updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		newLesson.LessonType,
		newLesson.Id,
	).Scan(&newLesson.CreatedAt, &newLesson.UpdatedAt)
	if err != nil {
		log.Println("Eror updating lesson in postgres method", err.Error())
		return nil, err
	}

	return newLesson, nil
}

// This function delete lesson info from postgres
func (s *lessonRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		lessons
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
		AND deleted_at IS NULL
	`
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	rowEffect, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect == 0 {
		log.Println("Nothing deleted, Lesson")
		return false, nil
	}

	return true, nil
}

// This function gets lesson in postgres
func (s *lessonRepo) Get(ctx context.Context, id string) (*repo.Lesson, error) {
	query := `
	SELECT 
		lessons.id,
		lessons.lesson_type,
    	topics.id,
    	topics.name,
		lessons.created_at,
		lessons.updated_at
	FROM 
		lessons
	JOIN 
		topics ON lessons.topic_id = topics.id
	JOIN
		levels ON topics.level_id = levels.id
	JOIN
		languages ON levels.language_id = languages.id
	WHERE lessons.id = $1
	AND lessons.deleted_at IS NULL 
	AND	topics.deleted_at IS NULL
	AND levels.deleted_at IS NULL
	AND languages.deleted_at IS NULL
	`

	var responseLesson repo.Lesson
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseLesson.Id,
		&responseLesson.LessonType,
		&responseLesson.TopicId,
		&responseLesson.TopicName,
		&responseLesson.CreatedAt,
		&responseLesson.UpdatedAt,
	)
	if err != nil {
		log.Println("Eror getting lesson in postgres method", err.Error())
		return nil, err
	}

	return &responseLesson, nil
}

// This function get all lesson with page and limit posgtres
func (s *lessonRepo) GetAll(ctx context.Context, lesson_id string) ([]*repo.Lesson, error) {
	fmt.Println("asdsadsad: ", lesson_id)
	query := `
	SELECT 
		lessons.id,
		lessons.lesson_type,
		topics.id,
		topics.name,
		lessons.created_at,
		lessons.updated_at
	FROM 
		lessons
	JOIN 
		topics ON lessons.topic_id = topics.id
	JOIN
		levels ON topics.level_id = levels.id
	JOIN
		languages ON levels.language_id = languages.id
	WHERE
		topic_id = $1
		AND lessons.deleted_at IS NULL
		AND	topics.deleted_at IS NULL
		AND levels.deleted_at IS NULL
		AND languages.deleted_at IS NULL
	`

	rows, err := s.db.QueryContext(ctx, query, lesson_id)
	if err != nil {
		fmt.Println("\x1b[32m 1", err,"\x1b[0m")

		log.Println("Error selecting lessons with page and limit in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseLessons []*repo.Lesson
	for rows.Next() {
		var lesson repo.Lesson
		err = rows.Scan(
			&lesson.Id,
			&lesson.LessonType,
			&lesson.TopicId,
			&lesson.TopicName,
			&lesson.CreatedAt,
			&lesson.UpdatedAt,
		)
		if err != nil {
			fmt.Println("\x1b[32m 2", err,"\x1b[0m")
			log.Println("Error scanning lesson in getall lesson method of postgres", err.Error())
			return nil, err
		}

		responseLessons = append(responseLessons, &lesson)
	}

	return responseLessons, nil
}
