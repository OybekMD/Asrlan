package v1

import (
	"asrlan-monolight/api/helper/parsing"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"asrlan-monolight/api/models"
	"asrlan-monolight/storage/repo"

	"github.com/gin-gonic/gin"
)

// @Security      BearerAuth
// @Summary 	  Create Lesson
// @Description   This Api for creating a new lesson
// @Tags 		  lessons
// @Accept 		  json
// @Produce 	  json
// @Param 		  LessonCreate body models.LessonCreate true "LessonCreate Model"
// @Success 	  201 {object} models.LessonResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/lesson [POST]
func (h *handlerV1) CreateLesson(ctx *gin.Context) {
	var body models.LessonCreate

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

	response, err := h.storage.Lesson().Create(
		ctxTime,
		&repo.Lesson{
			LessonType: body.LessonType,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.LessonResponse{
		Id:         response.Id,
		LessonType: response.LessonType,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Lesson
// @Description   This Api for updating lesson
// @Tags 		  lessons
// @Accept 		  json
// @Produce 	  json
// @Param 		  LessonUpdate body models.LessonUpdate true "Update LessonUpdate Model"
// @Success 	  200 {object} models.LessonResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/lesson [PUT]
func (h *handlerV1) UpdateLesson(ctx *gin.Context) {
	var body models.LessonUpdate

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

	lessonModel := &repo.Lesson{}
	err = parsing.StructToStruct(&body, lessonModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Lesson().Update(ctxTime, lessonModel)
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
// @Summary 	  Delete Lesson
// @Description   This Api for deleting lesson
// @Tags 		  lessons
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/lesson/{id} [DELETE]
func (h *handlerV1) DeleteLesson(ctx *gin.Context) {
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

	response, err := h.storage.Lesson().Delete(ctxTime, id)
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
// @Summary 	  Get Lesson
// @Description   This Api for get lesson
// @Tags 		  lessons
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.LessonResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/lesson/{id} [GET]
func (h *handlerV1) GetLesson(ctx *gin.Context) {
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

	lesson, err := h.storage.Lesson().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get lesson", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, lesson)
}

// @Security      BearerAuth
// @Summary       ListLessons
// @Description   This Api for get all lessons
// @Tags          lessons
// @Accept        json
// @Produce       json
// @Param         lesson_id query int64 true "LessonId"
// @Success 	  200 {object} []models.LessonResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/lessons [GET]
func (h *handlerV1) ListLessons(ctx *gin.Context) {
	lesson_id := ctx.Query("lesson_id")
	fmt.Println("Mountan:", lesson_id)

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

	lessons, err := h.storage.Lesson().GetAll(ctxTime, lesson_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(lessons) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found lessons")
		return
	}

	ctx.JSON(http.StatusOK, models.LessonListResponse{
		Lessons: lessons,
	})
}
