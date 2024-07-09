package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type contentFileRepo struct {
	db *sqlx.DB
}

func NewContentFile(db *sqlx.DB) repo.ContentFileStorageI {
	return &contentFileRepo{
		db: db,
	}
}

func QueryBuildCreateContentFile(reqData *repo.ContentFile) (string, []interface{}) {
	args := []interface{}{}
	columns := []string{}
	placeholders := []string{}

	addColumn := func(name string, value interface{}) {
		if value != "" {
			columns = append(columns, name)
			args = append(args, value)
			placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))
		}
	}
	columns = append(columns, "id")
	args = append(args, reqData.Id)
	columns = append(columns, "content_id")
	args = append(args, reqData.ContentId)

	addColumn("sound_data", reqData.SoundData)
	addColumn("image_data", reqData.ImageData)
	addColumn("video_data", reqData.VideoData)

	placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))

	query := fmt.Sprintf("INSERT INTO contentfiles (%s) VALUES (%s) RETURNING id, created_at",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}

// This function create contentFile in postgres
func (s *contentFileRepo) Create(ctx context.Context, contentFile *repo.ContentFile) (*repo.ContentFile, error) {
	query, args := QueryBuildCreateContentFile(contentFile)

	err := s.db.QueryRowContext(
		ctx,
		query, args...).Scan(&contentFile.Id, &contentFile.CreatedAt)
	if err != nil {
		log.Println("Eror creating contentFile in postgres method", err.Error())
		return nil, err
	}

	return contentFile, nil
}

// This function delete contentFile info from postgres
func (s *contentFileRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		contentFiles
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
		log.Println("Nothing deleted, ContentFile")
		return false, nil
	}

	return true, nil
}

// This function gets contentFile in postgres
func (s *contentFileRepo) Get(ctx context.Context, id string) (*repo.ContentFile, error) {
	query := `
	SELECT
		id, 
		content_id, 
		sound_data,
		image_data, 
		video_data, 
		created_at, 
		updated_at
	FROM
		users
	WHERE
		id = $1
	AND
		deleted_at IS NULL
	`

	var responseContentFile repo.ContentFile
	var nullSoundData, nullImageData, nullVideoData sql.NullString
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseContentFile.Id,
		&responseContentFile.ContentId,
		&nullSoundData,
		&nullImageData,
		&nullVideoData,
		&responseContentFile.CreatedAt,
	)
	if err != nil {
		log.Println("Eror getting user in postgres method", err.Error())
		return nil, err
	}
	if nullSoundData.Valid {
		responseContentFile.SoundData = nullSoundData.String
	}
	if nullImageData.Valid {
		responseContentFile.ImageData = nullImageData.String
	}
	if nullVideoData.Valid {
		responseContentFile.VideoData = nullVideoData.String
	}

	return &responseContentFile, nil
}
