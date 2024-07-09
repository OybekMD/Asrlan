package v1

import (
	"asrlan-monolight/api/helper/parsing"
	"context"
	"log"
	"net/http"
	"time"

	"asrlan-monolight/api/models"
	"asrlan-monolight/storage/repo"

	"github.com/gin-gonic/gin"
)

// @Security      BearerAuth
// @Summary 	  Create Book
// @Description   This Api for creating a new book
// @Tags 		  books
// @Accept 		  json
// @Produce 	  json
// @Param 		  BookCreate body models.BookCreate true "BookCreate Model"
// @Success 	  201 {object} models.BookResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/book [POST]
func (h *handlerV1) CreateBook(ctx *gin.Context) {
	var body models.BookCreate

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.storage.Book().Create(
		ctxTime,
		&repo.Book{
			Name: body.Name,
			Picture: body.Picture,
			BookFile: body.BookFile,
			LevelId: body.LevelId,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.BookResponse{
		Id:        response.Id,
		Name:      response.Name,
		Picture:   response.Picture,
		BookFile:   response.BookFile,
		LevelId:   response.LevelId,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Book
// @Description   This Api for updating book
// @Tags 		  books
// @Accept 		  json
// @Produce 	  json
// @Param 		  BookUpdate body models.BookUpdate true "Update BookUpdate Model"
// @Success 	  200 {object} models.BookResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/book [PUT]
func (h *handlerV1) UpdateBook(ctx *gin.Context) {
	var body models.BookUpdate

	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("failed to bind json", err.Error())
		return
	}

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	bookModel := &repo.Book{}
	err = parsing.StructToStruct(&body, bookModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Book().Update(ctxTime, bookModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotUpdatedMessage,
		})
		log.Println("failed to update user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// @Security      BearerAuth
// @Summary 	  Delete Book
// @Description   This Api for deleting book
// @Tags 		  books
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/book/{id} [DELETE]
func (h *handlerV1) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	response, err := h.storage.Book().Delete(ctxTime, id)
	if err != nil || !response {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotDeletedMessage,
		})
		log.Println("failed to delete user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

// @Security      BearerAuth
// @Summary 	  Get Book
// @Description   This Api for get book
// @Tags 		  books
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.BookResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/book/{id} [GET]
func (h *handlerV1) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	book, err := h.storage.Book().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get book", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, book)
}

// @Security      BearerAuth
// @Summary       ListBooks
// @Description   This Api for get all books
// @Tags          books
// @Accept        json
// @Produce       json
// @Param         id path string true "Level Id"
// @Success 	  200 {object} []models.BookResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/books/{id} [GET]
func (h *handlerV1) ListBooks(ctx *gin.Context) {
	book_id := ctx.Param("id")

	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to parse timeout", err.Error())
		return
	}

	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	books, err := h.storage.Book().GetAll(ctxTime, book_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(books) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found books")
		return
	}

	ctx.JSON(http.StatusOK, models.BookListResponse{
		Books: books,
	})
}
