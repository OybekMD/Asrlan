package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type contentRepo struct {
	db *sqlx.DB
}

func NewContent(db *sqlx.DB) repo.ContentStorageI {
	return &contentRepo{
		db: db,
	}
}

func QueryBuildCreateContent(reqData *repo.Content) (string, []interface{}) {
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
	columns = append(columns, "lesson_id")
	args = append(args, reqData.LessonId)
	columns = append(columns, "gentype")
	args = append(args, reqData.Gentype)
	columns = append(columns, "title")
	args = append(args, reqData.Title)

	addColumn("question", reqData.Question)
	addColumn("text_data", reqData.TextData)
	addColumn("arr_text", reqData.ArrText)
	addColumn("correct_answer", reqData.CorrectAnswer)

	placeholders = append(placeholders, fmt.Sprintf("$%d", len(args)))

	query := fmt.Sprintf("INSERT INTO contents (%s) VALUES (%s) RETURNING id, created_at",
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}

func QueryBuildUpdateContent(reqData *repo.Content) (string, []interface{}) {
	args := []interface{}{}
	updates := []string{}

	addUpdate := func(name string, value interface{}) {
		if value != "" {
			updates = append(updates, fmt.Sprintf("%s = $%d", name, len(args)+1))
			args = append(args, value)
		}
	}

	addUpdate("lesson_id", reqData.LessonId)
	addUpdate("gentype", reqData.Gentype)
	addUpdate("title", reqData.Title)
	addUpdate("question", reqData.Question)
	addUpdate("text_data", reqData.TextData)
	addUpdate("arr_text", reqData.ArrText)
	addUpdate("correct_answer", reqData.CorrectAnswer)

	query := fmt.Sprintf("UPDATE contents SET %s WHERE id = $%d RETURNING created_at, updated_at",
		strings.Join(updates, ", "),
		len(args)+1,
	)
	args = append(args, reqData.Id)

	return query, args
}

func convertUint8ToString(byteSlice []uint8) []string {
	// Convert the byte slice to a string
	str := string(byteSlice)

	// Remove the curly braces from the string
	str = strings.Trim(str, "{}")

	// Split the string by comma to get individual elements
	elements := strings.Split(str, ",")

	// Trim spaces from each element
	for i, element := range elements {
		elements[i] = strings.TrimSpace(element)
	}

	return elements
}

// This function create content in postgres
func (s *contentRepo) Create(ctx context.Context, content *repo.Content) (*repo.Content, error) {
	query, args := QueryBuildCreateContent(content)

	err := s.db.QueryRowContext(
		ctx,
		query, args...).Scan(&content.Id, &content.CreatedAt)
	if err != nil {
		log.Println("Eror creating content in postgres method", err.Error())
		return nil, err
	}

	return content, nil
}

// This function update content info from postgres
func (s *contentRepo) Update(ctx context.Context, newContent *repo.Content) (*repo.Content, error) {
	query, args := QueryBuildUpdateContent(newContent)
	err := s.db.QueryRowContext(
		ctx,
		query, args...).Scan(&newContent.CreatedAt, &newContent.UpdatedAt)
	if err != nil {
		log.Println("Eror updating content in postgres method", err.Error())
		return nil, err
	}

	return newContent, nil
}

// This function delete content info from postgres
func (s *contentRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		contents
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
		log.Println("Nothing deleted, Content")
		return false, nil
	}

	return true, nil
}

// This function gets content in postgres
func (s *contentRepo) Get(ctx context.Context, id string) (*repo.Content, error) {
	var res repo.Content
	var nullQuestion, nullTextData sql.NullString
	var nullArrText []pq.StringArray
	var nullCorrectAnswer sql.NullInt64
	query := `
	SELECT
		id,
		lesson_id,
		gentype,
		title,
		question,
		text_data,
		arr_text,
		correct_answer,
		created_at,
		updated_at
	FROM 
		contents
	WHERE
		id=$1
	AND deleted_at IS NULL
	`
	err := s.db.QueryRow(query, id).Scan(
		&res.Id,
		&res.LessonId,
		&res.Gentype,
		&res.Title,
		&nullQuestion,
		&nullTextData,
		pq.Array(&nullArrText),
		&nullCorrectAnswer,
		&res.CreatedAt,
		&res.UpdatedAt)
	if err != nil {
		log.Println("Error ReadContentDB Scaning: ", err.Error())
		return nil, err
	}
	if nullQuestion.Valid {
		res.Question = nullQuestion.String
	}
	if nullTextData.Valid {
		res.TextData = nullTextData.String
	}
	if nullCorrectAnswer.Valid {
		res.CorrectAnswer = nullCorrectAnswer.Int64
	}
	queryfile := `
		SELECT
			id,
			content_id,
			sound_data,
			image_data,
			video_data,
			created_at,
			updated_at
		FROM
			content_files
		WHERE
			content_id=$1
		AND
			deleted_at IS NULL
		`
	rowsfile, err := s.db.Query(queryfile, res.Id)
	if err != nil {
		return nil, err
	}
	var allcontentfiles []*repo.ContentFile
	for rowsfile.Next() {
		var contentfile repo.ContentFile
		var nullSoundData, nullImageData, nullVideoData sql.NullString
		err := rowsfile.Scan(
			&contentfile.Id,
			&contentfile.ContentId,
			&nullSoundData,
			&nullImageData,
			&nullVideoData,
			&contentfile.CreatedAt,
			&contentfile.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if nullSoundData.Valid {
			contentfile.SoundData = nullSoundData.String
		}
		if nullImageData.Valid {
			contentfile.ImageData = nullImageData.String
		}
		if nullVideoData.Valid {
			contentfile.VideoData = nullVideoData.String
		}
		allcontentfiles = append(allcontentfiles, &contentfile)
	}
	res.Contentfiles = allcontentfiles

	return &res, nil
}

// This function get all content with page and limit posgtres
func (s *contentRepo) GetAll(ctx context.Context, id string) ([]*repo.Content, int64, error) {
	var allContent []*repo.Content
	query := `
	SELECT
		id,
		lesson_id,
		gentype,
		title,
		question,
		text_data,
		arr_text,
		correct_answer,
		created_at,
		updated_at
	FROM 
		contents
	WHERE
		lesson_id=$1
	AND 
		deleted_at IS NULL
	`

	rows, err := s.db.Query(query, id)
	if err != nil {
		fmt.Println("\033[31m", "ERR1:", err, "\033[0m")
		return nil, 0, err
	}
	for rows.Next() {
		var content repo.Content
		var nullQuestion, nullTextData sql.NullString
		var nullArrText []uint8
		var nullCorrectAnswer sql.NullInt64
		err := rows.Scan(
			&content.Id,
			&content.LessonId,
			&content.Gentype,
			&content.Title,
			&nullQuestion,
			&nullTextData,
			&nullArrText,
			&nullCorrectAnswer,
			&content.CreatedAt,
			&content.UpdatedAt,
		)
		if err != nil {
			fmt.Println("\033[31m", "ERR2:", err, "\033[0m")
			log.Println("Error ReadContentDB Scaning: ", err.Error())
			return nil, 0, err
		}
		resarrtext := convertUint8ToString(nullArrText)

		content.ArrText = resarrtext
		// content.ArrText = []string{string(nullArrText)}
		if nullQuestion.Valid {
			content.Question = nullQuestion.String
		}
		if nullTextData.Valid {
			content.TextData = nullTextData.String
		}
		if nullCorrectAnswer.Valid {
			content.CorrectAnswer = nullCorrectAnswer.Int64
		}
		queryfile := `
		SELECT
			id,
			content_id,
			sound_data,
			image_data,
			video_data,
			created_at,
			updated_at
		FROM
			content_files
		WHERE
			content_id=$1
		AND
			deleted_at IS NULL
		`
		rowsfile, err := s.db.Query(queryfile, content.Id)
		if err != nil {
			fmt.Println("\033[31m", "ERR3:", err, "\033[0m")
			return nil, 0, err
		}
		var allcontentfiles []*repo.ContentFile
		for rowsfile.Next() {
			var contentfile repo.ContentFile
			var nullSoundData, nullImageData, nullVideoData sql.NullString
			err := rowsfile.Scan(
				&contentfile.Id,
				&contentfile.ContentId,
				&nullSoundData,
				&nullImageData,
				&nullVideoData,
				&contentfile.CreatedAt,
				&contentfile.UpdatedAt,
			)
			if err != nil {
				fmt.Println("\033[31m", "ERR4:", err, "\033[0m")
				return nil, 0, err
			}
			if nullSoundData.Valid {
				contentfile.SoundData = nullSoundData.String
			}
			if nullImageData.Valid {
				contentfile.ImageData = nullImageData.String
			}
			if nullVideoData.Valid {
				contentfile.VideoData = nullVideoData.String
			}
			allcontentfiles = append(allcontentfiles, &contentfile)
		}
		content.Contentfiles = allcontentfiles
		allContent = append(allContent, &content)
	}

	count := len(allContent)
	return allContent, int64(count), nil
}
