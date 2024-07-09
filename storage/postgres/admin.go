package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"
	"github.com/jmoiron/sqlx"
)

type adminRepo struct {
	db *sqlx.DB
}

func NewAdmin(db *sqlx.DB) repo.AdminStorageI {
	return &adminRepo{
		db: db,
	}
}

// This function create admin in postgres
func (s *adminRepo) Create(ctx context.Context, email string) (bool, error) {
	query := `
	INSERT INTO admins (
		email,
	)
	VALUES ($1)`

	_, err := s.db.ExecContext(
		ctx,
		query,
		email,
	)
	if err != nil {
		log.Println("Eror creating admin in postgres method", err.Error())
		return false, err
	}

	return true, nil
}

// This function delete admin info from postgres
func (s *adminRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	DELETE
		admins
	WHERE
		id = $1
	`

	resutl, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting admin", err.Error())
		return false, err
	}

	rowEffect, err := resutl.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect == 0 {
		log.Println("Nothing updated, Admin")
		return false, nil
	}

	return true, nil
}

// This function gets admin in postgres
func (s *adminRepo) Get(ctx context.Context, email string) (bool, error) {
	query := `
	SELECT
		id,
		email,
		created_at
	FROM
		admins
	WHERE
		email = $1
	`
	resutl, err := s.db.ExecContext(ctx, query, email)
	if err != nil {
		log.Println("Error deleting admin", err.Error())
		return false, err
	}

	rowEffect, err := resutl.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect == 0 {
		log.Println("Nothing found, Admin")
		return false, nil
	}

	return true, nil
}

// This function get all admin with page and limit posgtres
func (s *adminRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.Admin, error) {
	query := `
	SELECT
		id, 
		name, 
		login_key, 
		password, 
		avatar, 
		created_at
	FROM
		admins
	WHERE
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting admins with page and limit in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseAdmins []*repo.Admin
	for rows.Next() {
		var admin repo.Admin
		err = rows.Scan(
			&admin.Id,
			&admin.Email,
			&admin.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning admin in getall admin method of postgres", err.Error())
			return nil, err
		}

		responseAdmins = append(responseAdmins, &admin)
	}

	return responseAdmins, nil
}
