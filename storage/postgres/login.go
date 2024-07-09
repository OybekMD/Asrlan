package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
)

type loginRepo struct {
	db *sqlx.DB
}

func NewLogin(db *sqlx.DB) repo.LoginStorageI {
	return &loginRepo{
		db: db,
	}
}

// This function checks who the login and password belong to
func (l *loginRepo) Login(ctx context.Context, login string) (*repo.LoginResponse, error) {
	fmt.Println(login)
	query := `
	SELECT
		u.id, 
		u.name, 
		u.username,
		u.bio, 
		u.birth_day, 
		u.email,
		u.password,
		u.avatar, 
		u.coint,
		u.score,
		ul.id AS language_status,
		lev.id AS level_status,
		u.created_at
	FROM
		users u
	LEFT JOIN
		user_language ul ON u.id = ul.user_id AND ul.status = TRUE
	LEFT JOIN
		user_level lev ON u.id = lev.user_id AND lev.status = TRUE
	WHERE
		email = $1 OR
		username = $1
	AND
		deleted_at IS NULL
	`

	var responseUser repo.LoginResponse
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	err := l.db.QueryRowContext(ctx, query, login).Scan(
		&responseUser.Id,
		&responseUser.Name,
		&responseUser.Username,
		&nullBio,
		&nullBirthDay,
		&responseUser.Email,
		&responseUser.Password,
		&nullAvatar,
		&responseUser.Coint,
		&responseUser.Score,
		&responseUser.LanguageId,
		&responseUser.LevelId,
		&responseUser.CreatedAt,
	)
	if err != nil {
		fmt.Println("error 1:", err)
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

	fmt.Println("ketti:", responseUser)

	return &responseUser, nil
}

// This function generate login for user and returns login and password
func (l *loginRepo) SavePassword(ctx context.Context, req *repo.LoginPassword) (*repo.LoginPassword, error) {
	insertQuery := `
	INSERT INTO login_passwords (
		user_id,
		role,
		password
	)
	VALUES ($1, $2, $3)
	RETURNING
		user_id,
		role,
		login,
		password
	`

	var response repo.LoginPassword
	err := l.db.QueryRowContext(ctx, insertQuery, req.UserId, req.Role, req.Password).Scan(
		&response.UserId,
		&response.Role,
		&response.Login,
		&response.Password,
	)
	if err != nil {

		log.Println("error saving password in postgres", err.Error())
		return nil, err
	}

	return &response, nil
}

func (l *loginRepo) ResetPassword(ctx context.Context, req *repo.ResetPassword) (*repo.LoginResponse, error) {
	pp.Println(req)
	query := `
	UPDATE
		users
	SET
		password = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		email = $2
	AND
		deleted_at IS NULL
	RETURNING
		id, 
		name, 
		username,
		bio, 
		birth_day, 
		email,
		password,
		avatar, 
		coint,
		score,
		created_at
	`

	var responseUser repo.LoginResponse
	var nullBio, nullBirthDay, nullAvatar sql.NullString
	err := l.db.QueryRowContext(
		ctx,
		query,
		req.NewPassword,
		req.Email,
	).Scan(
		&responseUser.Id,
		&responseUser.Name,
		&responseUser.Username,
		&nullBio,
		&nullBirthDay,
		&responseUser.Email,
		&responseUser.Password,
		&nullAvatar,
		&responseUser.Coint,
		&responseUser.Score,
		&responseUser.CreatedAt,
	)
	if err != nil {
		log.Println("Error updating password")
		return nil, err
	}

	return &responseUser, nil
}

// This function gets id and role of user by login
func (l *loginRepo) GetUserByLogin(ctx context.Context, login string) (id, role string, err error) {
	query := `
	SELECT
		user_id,
		role
	FROM
		login_passwords
	WHERE
		login = $1
	AND
		deleted_at IS NULL
	`

	if err := l.db.QueryRowContext(ctx, query, login).Scan(&id, &role); err != nil {
		log.Println("Error finding user by login", err.Error())
		return "", "", err
	}

	return id, role, nil
}

// This method save refresh for user
func (l *loginRepo) SaveRefresh(ctx context.Context, role, id, refresh string) (bool, error) {
	if role == "admin" {
		return true, nil
	}
	queryS := `
	UPDATE
		students
	SET
		refresh_token = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $2
	AND
		deleted_at IS NULL
	`
	queryT := `
	UPDATE
		teachers
	SET
		refresh_token = $1,
		updated_at = CURRENT_TIMESTAMP
	WHERE
		id = $2
	AND
		deleted_at IS NULL
	`

	var query string
	if role == "teacher" {
		query = queryT
	} else if role == "student" {
		query = queryS
	} else {
		log.Println("tablename not found for update refresh")
		return false, errors.New("wrong table name")
	}

	result, err := l.db.ExecContext(ctx, query, refresh, id)
	if err != nil {
		log.Println("Error updating refresh", err.Error())
		return false, err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting roweffect", err.Error())
		return false, err
	}
	if effect == 0 {
		log.Println("Nothing refresh updated")
		return false, errors.New("refresh not updated")
	}

	return true, nil
}
