package models

import "asrlan-monolight/storage/repo"

type (
	BookCreate struct {
		Name     string `json:"name"`
		Picture  string `json:"picture"`
		BookFile string `json:"book_file"`
		LevelId  int64  `json:"level_id"`
	}

	BookRequest struct {
		Id string `json:"id"`
	}

	BookResponse struct {
		Id        int64  `json:"id"`
		Name      string `json:"lesson_name"`
		Picture   string `json:"picture"`
		BookFile  string `json:"book_file"`
		LevelId   int64  `json:"level_id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	BookListResponse struct {
		Books []*repo.Book
	}

	BookUpdate struct {
		Id       int64  `json:"id"`
		Name     string `json:"name"`
		Picture  string `json:"picture"`
		BookFile string `json:"book_file"`
		LevelId  int64  `json:"level_id"`
	}
)
