package v1

import (
	"context"
	"log"
	"net/http"
	"time"

	"asrlan-monolight/api/models"
	"asrlan-monolight/storage/repo"

	"github.com/gin-gonic/gin"
)

// @Security      BearerAuth
// @Summary 	  Create Activity
// @Description   This Api for creating a new activity
// @Tags 		  activitys
// @Accept 		  json
// @Produce 	  json
// @Param 		  ActivityCreate body models.ActivityCreate true "ActivityCreate Model"
// @Success 	  201 {object} models.ActivityResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/activity [POST]
func (h *handlerV1) CreateActivity(ctx *gin.Context) {
	var body models.ActivityCreate

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

	response, err := h.storage.Activity().Create(
		ctxTime,
		&repo.Activity{
			Day:      body.Day,
			Score:    body.Score,
			LessonId: body.LessonId,
			UserId:   body.UserId,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	createUserLesson, err := h.storage.UserLesson().Create(ctx, &repo.UserLesson{
		Score: body.Score,
		LessonId: body.LessonId,
		UserId: body.UserId,
	})

	if !createUserLesson {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user_lesson", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.ActivityResponse{
		Id:       response.Id,
		Day:      response.Day,
		Score:      response.Score,
		LessonId: response.LessonId,
		UserId:   response.UserId,
	})
}

// @Security      BearerAuth
// @Summary       ListActivitys
// @Description   This Api for get all activitys
// @Tags          activitys
// @Accept        json
// @Produce       json
// @Success 	  200 {object} []models.ActivityListResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/activitys [GET]
func (h *handlerV1) GetAllGroupedByMonth(ctx *gin.Context) {
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

	activitys, err := h.storage.Activity().GetAllGroupedByMonth(ctxTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(activitys) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found activitys")
		return
	}

	ctx.JSON(http.StatusOK, activitys)
}

// @Security      BearerAuth
// @Summary       GetAllGroupedByChoice
// @Description   This Api for get all activitys
// @Tags          activitys
// @Accept        json
// @Produce       json
// @Success 	  200 {object} []models.ActivityListResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/activitysch [GET]
func (h *handlerV1) GetAllGroupedByChoice(ctx *gin.Context) {
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

	activitys, err := h.storage.Activity().GetAllGroupedByChoice(ctxTime, "week")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(activitys) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found activitys")
		return
	}

	ctx.JSON(http.StatusOK, activitys)
}
