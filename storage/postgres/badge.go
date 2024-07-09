package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type badgeRepo struct {
	db *sqlx.DB
}

func NewBadge(db *sqlx.DB) repo.BadgeStorageI {
	return &badgeRepo{
		db: db,
	}
}

// This function create badge in postgres
func (s *badgeRepo) Create(ctx context.Context, badge *repo.Badge) (*repo.Badge, error) {
	query := `
	INSERT INTO badges (
		name,
		badge_date,
		badge_type,
		picture
	)
	VALUES ($1, $2, $3, $4) RETURNING id`

	err := s.db.QueryRowContext(
		ctx,
		query,
		badge.Name,
		badge.BadgeDate,
		badge.BadgeType,
		badge.Picture).Scan(&badge.Id)
	if err != nil {
		log.Println("Eror creating badge in postgres method", err.Error())
		return nil, err
	}

	return badge, nil
}

// This function update badge info from postgres
func (s *badgeRepo) Update(ctx context.Context, newBadge *repo.Badge) (*repo.Badge, error) {
	query := `
	UPDATE
		badges
	SET
		name = $1,
		badge_date = $2,
		badge_type = $3,
		picture = $4
	WHERE
		id = $5
	`
	_, err := s.db.ExecContext(
		ctx,
		query,
		newBadge.Name,
		newBadge.BadgeDate,
		newBadge.BadgeType,
		newBadge.Picture,
		newBadge.Id,
	)
	if err != nil {
		log.Println("Eror updating badge in postgres method", err.Error())
		return nil, err
	}

	return newBadge, nil
}

// This function delete badge info from postgres
func (s *badgeRepo) Delete(ctx context.Context, id string) (bool, error) {
	query_all_user := `
		DELETE FROM 
			user_badge
		WHERE
			badge_id = $1
	`
	if _, err := s.db.QueryContext(ctx, query_all_user, id); err != nil {
		return false, err
	}

	query := `
		DELETE FROM
			badges
		WHERE
			id = $1
	`

	if _, err := s.db.QueryContext(ctx, query, id); err != nil {
		return false, err
	}

	return true, nil
}

// This function gets badge in postgres
func (s *badgeRepo) Get(ctx context.Context, id string) (*repo.Badge, error) {
	query := `
	SELECT
		id, 
		name, 
		badge_date, 
		badge_type, 
		picture
	FROM
		badges
	WHERE
		id = $1
	`

	var responseBadge repo.Badge
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseBadge.Id,
		&responseBadge.Name,
		&responseBadge.BadgeDate,
		&responseBadge.BadgeType,
		&responseBadge.Picture,
	)
	if err != nil {
		log.Println("Eror getting badge in postgres method", err.Error())
		return nil, err
	}

	return &responseBadge, nil
}

// This function get all badge with page and limit posgtres
func (s *badgeRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.Badge, int64, error) {
	query := `
	SELECT
		id, 
		name, 
		badge_date, 
		badge_type, 
		picture,
		(SELECT COUNT(*) FROM badges) AS total_count
	FROM
		badges
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting badges with page and limit in postgres", err.Error())
		return nil, 0, err
	}
	defer rows.Close()

	var responseBadges []*repo.Badge
	var badgeCount int64
	for rows.Next() {
		var badge repo.Badge
		err = rows.Scan(
			&badge.Id,
			&badge.Name,
			&badge.BadgeDate,
			&badge.BadgeType,
			&badge.Picture,
			&badgeCount,
		)
		if err != nil {
			log.Println("Error scanning badge in getall badge method of postgres", err.Error())
			return nil, 0, err
		}

		responseBadges = append(responseBadges, &badge)
	}

	return responseBadges, badgeCount, nil
}
