package v1

import (
	"asrlan-monolight/api/helper/parsing"
	"asrlan-monolight/api/helper/utils"
	"context"
	"log"
	"net/http"
	"time"

	"asrlan-monolight/api/models"
	"asrlan-monolight/storage/repo"

	"github.com/gin-gonic/gin"
)

// @Security      BearerAuth
// @Summary 	  Create Topic
// @Description   This Api for creating a new topic
// @Tags 		  topics
// @Accept 		  json
// @Produce 	  json
// @Param 		  TopicCreate body models.TopicCreate true "TopicCreate Model"
// @Success 	  201 {object} models.TopicResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/topic [POST]
func (h *handlerV1) CreateTopic(ctx *gin.Context) {
	var body models.TopicCreate

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

	response, err := h.storage.Topic().Create(
		ctxTime,
		&repo.Topic{
			Name: body.Name,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.TopicResponse{
		Id:        response.Id,
		Name:      response.Name,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Topic
// @Description   This Api for updating topic
// @Tags 		  topics
// @Accept 		  json
// @Produce 	  json
// @Param 		  TopicUpdate body models.TopicUpdate true "Update TopicUpdate Model"
// @Success 	  200 {object} models.TopicResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/topic [PUT]
func (h *handlerV1) UpdateTopic(ctx *gin.Context) {
	var body models.TopicUpdate

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

	topicModel := &repo.Topic{}
	err = parsing.StructToStruct(&body, topicModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Topic().Update(ctxTime, topicModel)
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
// @Summary 	  Delete Topic
// @Description   This Api for deleting topic
// @Tags 		  topics
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/topic/{id} [DELETE]
func (h *handlerV1) DeleteTopic(ctx *gin.Context) {
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

	response, err := h.storage.Topic().Delete(ctxTime, id)
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
// @Summary 	  Get Topic
// @Description   This Api for get topic
// @Tags 		  topics
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.TopicResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/topic/{id} [GET]
func (h *handlerV1) GetTopic(ctx *gin.Context) {
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

	topic, err := h.storage.Topic().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get topic", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, topic)
}

// @Security      BearerAuth
// @Summary       ListTopics
// @Description   This Api for get all topics
// @Tags          topics
// @Accept        json
// @Produce       json
// @Param         page query uint64 true "Page"
// @Param         limit query uint64 true "Limit"
// @Success 	  200 {object} []models.TopicResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/topics [GET]
func (h *handlerV1) ListTopics(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to parse query params json" + errStr[0])
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

	topics, count, err := h.storage.Topic().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(topics) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found topics")
		return
	}

	ctx.JSON(http.StatusOK, models.TopicListResponse{
		Topics: topics,
		Count:  count,
	})
}
