package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type userLevelRepo struct {
	db *sqlx.DB
}

func NewUserLevel(db *sqlx.DB) repo.UserLevelStorageI {
	return &userLevelRepo{
		db: db,
	}
}

// This function create userLevel in postgres
func (s *userLevelRepo) Create(ctx context.Context, userLevel *repo.UserLevel) (bool, error) {
	queryStatus := `
	UPDATE
		user_level
	SET
		status = FALSE
	WHERE
		status = TRUE AND
		user_id = $1
		`
	_, err := s.db.ExecContext(
		ctx,
		queryStatus,
		userLevel.UserId,
	)

	if err != nil {
		log.Println("Error Replasing statius from false to true userLevel in postgres", err.Error())
		return false, err
	}

	query := `
	INSERT INTO user_level (
		user_id,
		level_id
	)
	VALUES ($1, $2)`

	_, err = s.db.ExecContext(
		ctx,
		query,
		userLevel.UserId,
		userLevel.LevelId,
	)

	if err != nil {
		log.Println("Eror creating userLevel in postgres method", err.Error())
		return false, err
	}

	return true, nil
}

// This function get all userLevel with page and limit posgtres
func (s *userLevelRepo) GetActive(ctx context.Context, userId string) (*repo.UserLevel, error) {
	query := `
	SELECT
		id,
		score,
		status,
		user_id,
		level_id,
		created_at
	FROM
		user_language
	WHERE 
		status = TRUE
		AND user_id = $1
	`

	var responseUserLevel repo.UserLevel
	err := s.db.QueryRowContext(ctx, query, userId).Scan(
		&responseUserLevel.Id,
		&responseUserLevel.Score,
		&responseUserLevel.Status,
		&responseUserLevel.UserId,
		&responseUserLevel.LevelId,
		&responseUserLevel.CreatedAt,
	)
	if err != nil {
		log.Println("Eror getting user_language in postgres method", err.Error())
		return nil, err
	}

	return &responseUserLevel, nil
}

func (s *userLevelRepo) GetAll(ctx context.Context, userId string) ([]*repo.UserLevel, error) {
	query := `
    SELECT
		id,
		score,
		status,
		user_id,
		level_id,
		created_at
    FROM
		user_language
    WHERE
		user_id = $1
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		log.Println("Error selecting userUserLevels for year in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseUserLevels []*repo.UserLevel
	for rows.Next() {
		var userLevel repo.UserLevel
		err = rows.Scan(
			&userLevel.Id,
			&userLevel.Score,
			&userLevel.Status,
			&userLevel.UserId,
			&userLevel.LevelId,
			&userLevel.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning userLevel in YearlyUsersUserLevelsByYear method of postgres", err.Error())
			return nil, err
		}

		responseUserLevels = append(responseUserLevels, &userLevel)
	}

	return responseUserLevels, nil
}
