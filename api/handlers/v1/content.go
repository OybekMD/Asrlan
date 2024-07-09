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
// @Summary 	  Create Content
// @Description   This Api for creating a new content
// @Tags 		  contents
// @Accept 		  json
// @Produce 	  json
// @Param 		  ContentCreate body models.ContentCreate true "ContentCreate Model"
// @Success 	  201 {object} models.ContentResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/content [POST]
func (h *handlerV1) CreateContent(ctx *gin.Context) {
	var body models.ContentCreate

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

	response, err := h.storage.Content().Create(
		ctxTime,
		&repo.Content{
			LessonId:      body.LessonId,
			Gentype:       body.Gentype,
			Title:         body.Title,
			Question:      body.Question,
			TextData:      body.TextData,
			ArrText:       body.ArrText,
			CorrectAnswer: body.CorrectAnswer,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.ContentResponse{
		Id:            response.Id,
		LessonId:      response.LessonId,
		Gentype:       response.Gentype,
		Title:         response.Title,
		Question:      response.Question,
		TextData:      response.TextData,
		ArrText:       response.ArrText,
		CorrectAnswer: response.CorrectAnswer,
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Content
// @Description   This Api for updating content
// @Tags 		  contents
// @Accept 		  json
// @Produce 	  json
// @Param 		  ContentUpdate body models.ContentUpdate true "Update ContentUpdate Model"
// @Success 	  200 {object} models.ContentResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/content [PUT]
func (h *handlerV1) UpdateContent(ctx *gin.Context) {
	var body models.ContentUpdate

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

	contentModel := &repo.Content{}
	err = parsing.StructToStruct(&body, contentModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Content().Update(ctxTime, contentModel)
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
// @Summary 	  Delete Content
// @Description   This Api for deleting content
// @Tags 		  contents
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/content/{id} [DELETE]
func (h *handlerV1) DeleteContent(ctx *gin.Context) {
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

	response, err := h.storage.Content().Delete(ctxTime, id)
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
// @Summary 	  Get Content
// @Description   This Api for get content
// @Tags 		  contents
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.ContentResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/content/{id} [GET]
func (h *handlerV1) GetContent(ctx *gin.Context) {
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

	content, err := h.storage.Content().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get content", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, content)
}

// @Security      BearerAuth
// @Summary       ListContents
// @Description   This Api for get all contents
// @Tags          contents
// @Accept        json
// @Produce       json
// @Param		  id  query	  string	true	"ID"
// @Success 	  200 {object} []models.ContentResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/contents [GET]
func (h *handlerV1) ListContents(ctx *gin.Context) {
	id := ctx.Query("id")

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

	contents, count, err := h.storage.Content().GetAll(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(contents) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found contents")
		return
	}

	ctx.JSON(http.StatusOK, models.ContentListResponse{
		Contents: contents,
		Count:    count,
	})
}
