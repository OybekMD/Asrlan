package postgres

import (
	"asrlan-monolight/storage/repo"
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp/v3"
)

type bookRepo struct {
	db *sqlx.DB
}

func NewBook(db *sqlx.DB) repo.BookStorageI {
	return &bookRepo{
		db: db,
	}
}

// This function create book in postgres
func (s *bookRepo) Create(ctx context.Context, book *repo.Book) (*repo.Book, error) {
	query := `
	INSERT INTO books(
		name,
		picture,
		book_file,
		level_id
	)
	VALUES ($1, $2, $3, $4) 
	RETURNING 
		id,
		created_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		book.Name,
		book.Picture,
		book.BookFile,
		book.LevelId,
	).Scan(&book.Id, &book.CreatedAt)
	if err != nil {
		log.Println("Eror creating book in postgres method", err.Error())
		return nil, err
	}

	return book, nil
}

// This function update book info from postgres
func (s *bookRepo) Update(ctx context.Context, newBook *repo.Book) (*repo.Book, error) {
	query := `
	UPDATE
		books
	SET
		name=$1,
		picture=$2,
		book_file=$3,
		level_id=$4,
		updated_at=CURRENT_TIMESTAMP
	WHERE
		id=$5
	AND deleted_at IS NULL
	RETURNING
		created_at,
		updated_at
	`
	err := s.db.QueryRowContext(
		ctx,
		query,
		newBook.Name,
		newBook.Picture,
		newBook.BookFile,
		newBook.LevelId,
		newBook.Id,
	).Scan(&newBook.CreatedAt, &newBook.UpdatedAt)
	if err != nil {
		log.Println("Eror updating book in postgres method", err.Error())
		return nil, err
	}

	return newBook, nil
}

// This function delete book info from postgres
func (s *bookRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `
	UPDATE
		books
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
		log.Println("Nothing deleted, Book")
		return false, nil
	}

	return true, nil
}

// This function gets book in postgres
func (s *bookRepo) Get(ctx context.Context, id string) (*repo.Book, error) {
	query := `
	SELECT 
		id,
    	name,
    	picture,
    	book_file,
    	level_id,
		created_at,
		updated_at
	FROM 
		books
	WHERE 
		id = $1
		AND deleted_at IS NULL 
	`

	var responseBook repo.Book
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&responseBook.Id,
		&responseBook.Name,
		&responseBook.Picture,
		&responseBook.BookFile,
		&responseBook.LevelId,
		&responseBook.CreatedAt,
		&responseBook.UpdatedAt,
	)
	if err != nil {
		log.Println("Eror getting book in postgres method", err.Error())
		return nil, err
	}

	return &responseBook, nil
}

// This function get all book with page and limit posgtres
func (s *bookRepo) GetAll(ctx context.Context, book_id string) ([]*repo.Book, error) {
	pp.Println(book_id)
	query := `
	SELECT 
		id,
    	name,
    	picture,
    	book_file,
    	level_id,
		created_at,
		updated_at
	FROM 
		books
	WHERE 
		level_id = $1
		AND deleted_at IS NULL
	`

	rows, err := s.db.QueryContext(ctx, query, book_id)
	if err != nil {
		log.Println("Error selecting books with page and limit in postgres", err.Error())
		return nil, err
	}
	defer rows.Close()

	var responseBooks []*repo.Book
	for rows.Next() {
		var book repo.Book
		err = rows.Scan(
			&book.Id,
			&book.Name,
			&book.Picture,
			&book.BookFile,
			&book.LevelId,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning book in getall book method of postgres", err.Error())
			return nil, err
		}

		responseBooks = append(responseBooks, &book)
	}

	return responseBooks, nil
}
