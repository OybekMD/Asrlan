package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	hash "asrlan-monolight/api/helper/hashing"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

// This function create user in postgres
func (s *userRepo) Create(ctx context.Context, user *repo.User) (*repo.User, error) {
	query := `
	INSERT INTO users (
		id,
		name,
		username,
		email,
		password,
		refresh_token
	)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING
		id,
		name,
		username,
		email,
		password,
		coint,
		score,
		refresh_token,
		created_at
	`

	var respUser repo.User
	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Id,
		user.Name,
		user.Username,
		user.Email,
		user.Password,
		user.RefreshToken,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.Username,
		&respUser.Email,
		&respUser.Password,
		&respUser.Coint,
		&respUser.Score,
		&respUser.RefreshToken,
		&respUser.CreatedAt,
	)
	if err != nil {
		log.Println("Eror creating user in postgres method", err.Error())
		return nil, err
	}

	return &respUser, nil
}

// This function update user info from postgres
func (s *userRepo) Update(ctx context.Context, newUser *repo.User) (*repo.User, error) {
	query := `
	UPDATE
		users
	SET
		name = $1,
		username = $2,
		bio = $3,
		birth_day = $4,
		avatar = $5,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $6
	AND
		deleted_at IS NULL
	RETURNING
		id,
		name,
		username,
		bio,
		birth_day,
		email,
		avatar,
		coint,
		score,
		created_at,
		updated_at
	`

	pp.Println(newUser)

	var updatedUser repo.User
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	err := s.db.QueryRowContext(
		ctx,
		query,
		newUser.Name,
		newUser.Username,
		newUser.Bio,
		newUser.BirthDay,
		newUser.Avatar,
		newUser.Id,
	).Scan(
		&updatedUser.Id,
		&updatedUser.Name,
		&updatedUser.Username,
		&nullBio,
		&nullBirthDay,
		&updatedUser.Email,
		&nullAvatar,
		&updatedUser.Coint,
		&updatedUser.Score,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		log.Println("Eror updating user in postgres method", err.Error())
		return nil, err
	}
	if nullBio.Valid {
		updatedUser.Bio = nullBio.String
	}
	if nullBirthDay.Valid {
		updatedUser.BirthDay = nullBirthDay.String
	}
	if nullAvatar.Valid {
		updatedUser.Avatar = nullAvatar.String
	}

	return &updatedUser, nil
}

func (s *userRepo) UpdatePassword(ctx context.Context, newPassword *repo.UserUpdatePassword) (bool, error) {
	pp.Println(newPassword)
	queryGet := `
	SELECT
		password
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	var getPassword string
	err := s.db.QueryRowContext(ctx, queryGet, newPassword.Id).Scan(
		&getPassword,
	)

	fmt.Println("\x1b[32m", getPassword,"\x1b[0m")


	if err != nil {
		log.Println("Eror updating user in postgres method", err.Error())
		return false, err
	}
	if !hash.CheckPasswordHash(newPassword.OldPassword, getPassword) {
		return false, nil
	}

	hashedPassword,err := hash.HashPassword(newPassword.NewPassword)
	if err != nil {
		return false, err
	}

	query := `
	UPDATE
		users
	SET
		password = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $2
	AND
		deleted_at IS NULL
	`
	_, err = s.db.ExecContext(
		ctx,
		query,
		hashedPassword,
		newPassword.Id,
	)
	if err != nil {
		log.Println("Eror updating user in postgres method", err)
		return false, err
	}

	return true, nil
}

// This function delete user info from postgres
func (s *userRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		users
	SET
		deleted_at = CURRENT_TIMESTAMP
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting user", err.Error())
		return false, err
	}

	rowEffect, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting row effect", err.Error())
		return false, err
	}

	if rowEffect == 0 {
		log.Println("Nothing deleted, User")
		return false, nil
	}

	return true, nil
}

// This function gets user in postgres
func (s *userRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	query := `
	SELECT
		id, 
		name, 
		username,
		bio, 
		birth_day, 
		email, 
		avatar, 
		coint,
		score,
		created_at
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	var responseUser repo.User
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseUser.Id,
		&responseUser.Name,
		&responseUser.Username,
		&nullBio,
		&nullBirthDay,
		&responseUser.Email,
		&nullAvatar,
		&responseUser.Coint,
		&responseUser.Score,
		&responseUser.CreatedAt,
	)
	if err != nil {
		log.Println("Eror getting user in postgres method", err.Error())
		return nil, err
	}
	if nullBio.Valid {
		responseUser.Bio = nullBio.String
	}
	if nullBirthDay.Valid {
		responseUser.BirthDay = nullBirthDay.String
	}
	if nullAvatar.Valid {
		responseUser.Avatar = nullAvatar.String
	}

	return &responseUser, nil
}

// This function get all user with page and limit posgtres
func (s *userRepo) GetAll(ctx context.Context, page, limit uint64) ([]*repo.User, error) {
	query := `
	SELECT
		id, 
		name, 
		username,
		bio, 
		birth_day, 
		email, 
		avatar, 
		coint,
		score,
		created_at
	FROM
		users
	WHERE
		deleted_at IS NULL
	LIMIT $1
	OFFSET $2
	`

	offset := limit * (page - 1)
	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error selecting users with page and limit in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseUsers []*repo.User
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	for rows.Next() {
		var user repo.User
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.Username,
			&user.Bio,
			&user.BirthDay,
			&user.Email,
			&user.Avatar,
			&user.Coint,
			&user.Score,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning user in getall user method of postgres", err.Error())
			return nil, err
		}
		if nullBio.Valid {
			user.Bio = nullBio.String
		}
		if nullBirthDay.Valid {
			user.BirthDay = nullBirthDay.String
		}
		if nullAvatar.Valid {
			user.Avatar = nullAvatar.String
		}

		responseUsers = append(responseUsers, &user)
	}

	return responseUsers, nil
}

func (s *userRepo) CheckField(ctx context.Context, field, value string) (bool, error) {

	query := fmt.Sprintf(
		`SELECT count(1) 
		FROM users WHERE %s = $1 
		AND deleted_at IS NULL`, field)

	var isExists int

	row := s.db.QueryRowContext(ctx, query, value)
	if err := row.Scan(&isExists); err != nil {
		return true, err
	}

	if isExists == 0 {
		return false, nil
	}

	return true, nil
}

func (s *userRepo) CheckUsername(ctx context.Context, id, username string) (bool, error) {
	query :=
		`
		SELECT
			id
		FROM 
			users
		WHERE username = $1
		AND deleted_at IS NULL`

	var nullpgetid sql.NullString
	var pgetid string

	row := s.db.QueryRowContext(ctx, query, username)
	err := row.Scan(&nullpgetid)
	if err != nil {
		return true, nil
	}

	if nullpgetid.Valid {
		pgetid = nullpgetid.String
	} else {
		return false, errors.New("null Valid is not validated")
	}

	if pgetid != id {
		return false, nil
	}

	return true, nil
}