package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type userLanguageRepo struct {
	db *sqlx.DB
}

func NewUserLanguage(db *sqlx.DB) repo.UserLanguageStorageI {
	return &userLanguageRepo{
		db: db,
	}
}

// This function create userLanguage in postgres
func (s *userLanguageRepo) Create(ctx context.Context, userLanguage *repo.UserLanguage) (bool, error) {
	queryStatus := `
	UPDATE
		user_language
	SET
		status = FALSE
	WHERE
		status = TRUE AND
		user_id = $1
		`
	_, err := s.db.ExecContext(
		ctx,
		queryStatus,
		userLanguage.UserId,
	)

	if err != nil {
		log.Println("Error Replasing statius from false to true userLanguage in postgres", err.Error())
		return false, err
	}

	query := `
	INSERT INTO user_language (
		user_id,
		language_id
	)
	VALUES ($1, $2)`

	_, err = s.db.ExecContext(
		ctx,
		query,
		userLanguage.UserId,
		userLanguage.LanguageId,
	)

	if err != nil {
		log.Println("Eror creating userLanguage in postgres method", err.Error())
		return false, err
	}

	return true, nil
}

// This function get all userLanguage with page and limit posgtres
func (s *userLanguageRepo) GetActive(ctx context.Context, userId string) (*repo.UserLanguage, error) {
	query := `
	SELECT
		id,
		score,
		status,
		user_id,
		language_id,
		created_at
	FROM
		user_language
	WHERE 
		status = TRUE
		AND user_id = $1
	`

	var responseUserLanguage repo.UserLanguage
	err := s.db.QueryRowContext(ctx, query, userId).Scan(
		&responseUserLanguage.Id,
		&responseUserLanguage.Score,
		&responseUserLanguage.Status,
		&responseUserLanguage.UserId,
		&responseUserLanguage.LanguageId,
		&responseUserLanguage.CreatedAt,
	)
	if err != nil {
		log.Println("Eror getting user_language in postgres method", err.Error())
		return nil, err
	}

	return &responseUserLanguage, nil
}

func (s *userLanguageRepo) GetAll(ctx context.Context, userId string) ([]*repo.UserLanguage, error) {
	query := `
    SELECT
		id,
		score,
		status,
		user_id,
		language_id,
		created_at
    FROM
		user_language
    WHERE
		user_id = $1
    `

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		log.Println("Error selecting userUserLanguages for year in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseUserLanguages []*repo.UserLanguage
	for rows.Next() {
		var userLanguage repo.UserLanguage
		err = rows.Scan(
			&userLanguage.Id,
			&userLanguage.Score,
			&userLanguage.Status,
			&userLanguage.UserId,
			&userLanguage.LanguageId,
			&userLanguage.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning userLanguage in YearlyUsersUserLanguagesByYear method of postgres", err.Error())
			return nil, err
		}

		responseUserLanguages = append(responseUserLanguages, &userLanguage)
	}

	return responseUserLanguages, nil
}
