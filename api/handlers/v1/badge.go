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
// @Summary 	  Create Badge
// @Description   This Api for creating a new badge
// @Tags 		  badges
// @Accept 		  json
// @Produce 	  json
// @Param 		  BadgeCreate body models.BadgeCreate true "BadgeCreate Model"
// @Success 	  201 {object} models.BadgeResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/badge [POST]
func (h *handlerV1) CreateBadge(ctx *gin.Context) {
	var body models.BadgeCreate

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

	response, err := h.storage.Badge().Create(
		ctxTime,
		&repo.Badge{
			Name:      body.Name,
			BadgeDate: body.BadgeDate,
			BadgeType: body.BadgeType,
			Picture:   body.Picture,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.BadgeResponse{
		Id:        response.Id,
		Name:      response.Name,
		BadgeDate: response.BadgeDate,
		BadgeType: response.BadgeType,
		Picture:   response.Picture,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Badge
// @Description   This Api for updating badge
// @Tags 		  badges
// @Accept 		  json
// @Produce 	  json
// @Param 		  BadgeUpdate body models.BadgeUpdate true "Update BadgeUpdate Model"
// @Success 	  200 {object} models.BadgeResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/badge [PUT]
func (h *handlerV1) UpdateBadge(ctx *gin.Context) {
	var body models.BadgeUpdate

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

	badgeModel := &repo.Badge{}
	err = parsing.StructToStruct(&body, badgeModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Badge().Update(ctxTime, badgeModel)
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
// @Summary 	  Delete Badge
// @Description   This Api for deleting badge
// @Tags 		  badges
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/badge/{id} [DELETE]
func (h *handlerV1) DeleteBadge(ctx *gin.Context) {
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

	response, err := h.storage.Badge().Delete(ctxTime, id)
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
// @Summary 	  Get Badge
// @Description   This Api for get badge
// @Tags 		  badges
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.BadgeResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/badge/{id} [GET]
func (h *handlerV1) GetBadge(ctx *gin.Context) {
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

	badge, err := h.storage.Badge().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get badge", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, badge)
}

// @Security      BearerAuth
// @Summary       ListBadges
// @Description   This Api for get all badges
// @Tags          badges
// @Accept        json
// @Produce       json
// @Param         page query uint64 true "Page"
// @Param         limit query uint64 true "Limit"
// @Success 	  200 {object} []models.BadgeResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/badges [GET]
func (h *handlerV1) ListBadges(ctx *gin.Context) {
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

	badges, count, err := h.storage.Badge().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(badges) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found badges")
		return
	}

	ctx.JSON(http.StatusOK, models.BadgeListResponse{
		Badges: badges,
		Count:  count,
	})
}
