package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type topicRepo struct {
	db *sqlx.DB
}

func NewTopic(db *sqlx.DB) repo.TopicStorageI {
	return &topicRepo{
		db: db,
	}
}

// This function create topic in postgres
func (s *topicRepo) Create(ctx context.Context, topic *repo.Topic) (*repo.Topic, error) {
	query := `
	INSERT INTO topics(
		name,
		level_id
	)
	VALUES ($1, $2) 
	RETURNING 
		id,
		created_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		topic.Name,
		topic.LevelId).Scan(&topic.Id, &topic.CreatedAt)
	if err != nil {
		log.Println("Eror creating topic in postgres method", err.Error())
		return nil, err
	}

	return topic, nil
}

// This function update topic info from postgres
func (s *topicRepo) Update(ctx context.Context, newTopic *repo.Topic) (*repo.Topic, error) {
	query := `
	UPDATE
		topics
	SET
		name=$1,
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
		newTopic.Name,
		newTopic.Id,
	).Scan(&newTopic.CreatedAt, &newTopic.UpdatedAt)
	if err != nil {
		log.Println("Eror updating topic in postgres method", err.Error())
		return nil, err
	}

	return newTopic, nil
}

// This function delete topic info from postgres
func (s *topicRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		topics
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
		log.Println("Nothing deleted, Topic")
		return false, nil
	}

	return true, nil
}

// This function gets topic in postgres
func (s *topicRepo) Get(ctx context.Context, id string) (*repo.Topic, error) {
	query := `
	SELECT 
		topics.id,
    	topics.name,
    	levels.id,
    	levels.name,
		topics.created_at,
		topics.updated_at
	FROM 
		topics
	JOIN 
		levels ON topics.level_id = levels.id
	JOIN
		languages ON levels.language_id = languages.id
	WHERE topics.id = $1 
	AND topics.deleted_at IS NULL 
	AND	levels.deleted_at IS NULL
	AND topics.deleted_at IS NULL
	`

	var responseTopic repo.Topic
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseTopic.Id,
		&responseTopic.Name,
		&responseTopic.LevelId,
		&responseTopic.LevelName,
		&responseTopic.CreatedAt,
		&responseTopic.UpdatedAt,
	)
	if err != nil {
		log.Println("Eror getting topic in postgres method", err.Error())
		return nil, err
	}

	return &responseTopic, nil
}

// This function get all topic with page and limit posgtres
func (s *topicRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.Topic, int64, error) {
	query := `
	SELECT 
		topics.id,
		topics.name,
		levels.id,
		levels.name,
		topics.created_at,
		topics.updated_at
	FROM 
		topics
	JOIN
		levels ON topics.level_id = levels.id
	JOIN
		languages ON levels.language_id = languages.id
	WHERE
		topics.deleted_at IS NULL
		AND	levels.deleted_at IS NULL
		AND languages.deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting topics with page and limit in postgres", err.Error())
		return nil, 0, err
	}
	defer rows.Close()

	var responseTopics []*repo.Topic
	for rows.Next() {
		var topic repo.Topic
		err = rows.Scan(
			&topic.Id,
			&topic.Name,
			&topic.LevelId,
			&topic.LevelName,
			&topic.CreatedAt,
			&topic.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning topic in getall topic method of postgres", err.Error())
			return nil, 0, err
		}

		responseTopics = append(responseTopics, &topic)
	}

	count := len(responseTopics)

	return responseTopics, int64(count), nil
}
