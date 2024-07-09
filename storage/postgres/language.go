package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type languageRepo struct {
	db *sqlx.DB
}

func NewLanguage(db *sqlx.DB) repo.LanguageStorageI {
	return &languageRepo{
		db: db,
	}
}

// This function create language in postgres
func (s *languageRepo) Create(ctx context.Context, language *repo.Language) (*repo.Language, error) {
	query := `
	INSERT INTO languages(
		name,
		picture
	)
	VALUES ($1, $2) 
	RETURNING 
		id,
		created_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		language.Name,
		language.Picture).Scan(&language.Id, &language.CreatedAt)
	if err != nil {
		log.Println("Eror creating language in postgres method", err.Error())
		return nil, err
	}

	return language, nil
}

// This function update language info from postgres
func (s *languageRepo) Update(ctx context.Context, newLanguage *repo.Language) (*repo.Language, error) {
	query := `
	UPDATE
		languages
	SET
		name=$1,
		picture=$2,
		updated_at=CURRENT_TIMESTAMP
	WHERE
		id=$3
	AND deleted_at IS NULL
	RETURNING
		created_at,
		updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		newLanguage.Name,
		newLanguage.Picture,
		newLanguage.Id,
	).Scan(&newLanguage.CreatedAt, &newLanguage.UpdatedAt)
	if err != nil {
		log.Println("Eror updating language in postgres method", err.Error())
		return nil, err
	}

	return newLanguage, nil
}

// This function delete language info from postgres
func (s *languageRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		languages
	SET
		deleted_at=CURRENT_TIMESTAMP
	WHERE
		id=$1
	`
	if _, err := s.db.QueryContext(ctx, query, id); err != nil {
		return false, err
	}
	return false, nil
}

// This function gets language in postgres
func (s *languageRepo) Get(ctx context.Context, id string) (*repo.Language, error) {
	query := `
	SELECT
		id,
		name,
		picture,
		created_at,
		updated_at
	FROM 
		languages
	WHERE
		id=$1
	AND deleted_at IS NULL
	`

	var responseLanguage repo.Language
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseLanguage.Id,
		&responseLanguage.Name,
		&responseLanguage.Picture,
		&responseLanguage.CreatedAt,
		&responseLanguage.UpdatedAt,
	)
	if err != nil {
		log.Println("Eror getting language in postgres method", err.Error())
		return nil, err
	}

	return &responseLanguage, nil
}

// This function get all language with page and limit posgtres
func (s *languageRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.Language, int64, error) {
	query := `
	SELECT
		id, 
		name,
		picture,
		created_at,
		updated_at
	FROM 
		languages 
	WHERE 
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting languages with page and limit in postgres", err.Error())
		return nil, 0, err
	}
	defer rows.Close()

	var responseLanguages []*repo.Language
	for rows.Next() {
		var language repo.Language
		err = rows.Scan(
			&language.Id,
			&language.Name,
			&language.Picture,
			&language.CreatedAt,
			&language.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning language in getall language method of postgres", err.Error())
			return nil, 0, err
		}

		responseLanguages = append(responseLanguages, &language)
	}

	count := len(responseLanguages)

	return responseLanguages, int64(count), nil
}

func (s *languageRepo) GetAllForRegister(ctx context.Context) ([]*repo.RegisterLanguage, error) {
	query := `
	SELECT
		l.id,
		l.name,
		l.picture,
		COUNT(ul.id) AS learned_users_count
	FROM
		languages l
	LEFT JOIN
		user_language ul ON l.id = ul.language_id
	WHERE
		l.deleted_at IS NULL
	GROUP BY
		l.id
	`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("Error selecting languages with page and limit in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseLanguages []*repo.RegisterLanguage
	for rows.Next() {
		var language repo.RegisterLanguage
		err = rows.Scan(
			&language.Id,
			&language.Name,
			&language.Picture,
			&language.UserCount,
		)
		if err != nil {
			log.Println("Error scanning language in getall language method of postgres", err.Error())
			return nil, err
		}

		responseLanguages = append(responseLanguages, &language)
	}

	return responseLanguages, nil
}
