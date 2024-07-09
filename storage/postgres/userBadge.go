package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type userUserBadgeRepo struct {
	db *sqlx.DB
}

func NewUserBadge(db *sqlx.DB) repo.UserBadgeStorageI {
	return &userUserBadgeRepo{
		db: db,
	}
}

// This function create userBadge in postgres
func (s *userUserBadgeRepo) Create(ctx context.Context, userBadge *repo.UserBadge) (bool, error) {
	query := `
	INSERT INTO user_badge (
		user_id,
		badge_id
	)
	VALUES ($1, $2)`

	_, err := s.db.ExecContext(
		ctx,
		query,
		userBadge,
	)
	if err != nil {
		log.Println("Eror creating userBadge in postgres method", err.Error())
		return false, err
	}

	return true, nil
}

// This function delete userBadge info from postgres
func (s *userUserBadgeRepo) Delete(ctx context.Context, user_id, badge_id string) (bool, error) {
	query := `
		DELETE FROM
			user_badge
		WHERE
			user_id = $1 AND
			badge_id = $2
	`

	result, err := s.db.ExecContext(ctx, query, user_id, badge_id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}
	if rowsAffected == 0 {
		return false, nil
	}

	return true, nil
}

// This function get all userBadge with page and limit posgtres
func (s *userUserBadgeRepo) AllUsersBadgeByUserId(ctx context.Context, page, limit uint64) ([]*repo.Badge, int64, error) {
	query := `
	SELECT
		id,
		name,
		picture,
		(SELECT COUNT(*) FROM userUserBadges) AS total_count
	FROM
		badges
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting userUserBadges with page and limit in postgres", err.Error())
		return nil, 0, err
	}
	defer rows.Close()

	var responseUserBadges []*repo.Badge
	var userBadgeCount int64
	for rows.Next() {
		var userBadge repo.Badge
		err = rows.Scan(
			&userBadge.Id,
			&userBadge.Name,
			&userBadge.Picture,
			&userBadgeCount,
		)
		if err != nil {
			log.Println("Error scanning userBadge in getall userBadge method of postgres", err.Error())
			return nil, 0, err
		}

		responseUserBadges = append(responseUserBadges, &userBadge)
	}

	return responseUserBadges, userBadgeCount, nil
}

func (s *userUserBadgeRepo) YearlyUsersBadgesByYear(ctx context.Context, year int) ([]*repo.Badge, error) {
	query := `
    SELECT
        id, 
        name,
		badge_date
        picture
    FROM
        userUserBadges
    WHERE
        EXTRACT(YEAR FROM badge_date) = $1
		AND badge_type = 'month'
		ORDER BY
        EXTRACT(MONTH FROM badge_date) ASC
    `

	rows, err := s.db.QueryContext(ctx, query, year)
	if err != nil {
		log.Println("Error selecting userUserBadges for year in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseUserBadges []*repo.Badge
	for rows.Next() {
		var userBadge repo.Badge
		err = rows.Scan(
			&userBadge.Id,
			&userBadge.Name,
			&userBadge.BadgeDate,
			&userBadge.Picture,
		)
		if err != nil {
			log.Println("Error scanning userBadge in YearlyUsersBadgesByYear method of postgres", err.Error())
			return nil, err
		}

		responseUserBadges = append(responseUserBadges, &userBadge)
	}

	return responseUserBadges, nil
}
