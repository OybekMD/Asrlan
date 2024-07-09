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
// @Summary 	  Create UserBadge
// @Description   This Api for creating a new userBadge
// @Tags 		  userBadges
// @Accept 		  json
// @Produce 	  json
// @Param 		  UserBadgeCreate body models.UserBadgeCreate true "UserBadgeCreate Model"
// @Success 	  201 {object} bool
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/userbadge [POST]
func (h *handlerV1) CreateUserBadge(ctx *gin.Context) {
	var body models.UserBadgeCreate

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

	response, err := h.storage.UserBadge().Create(
		ctxTime,
		&repo.UserBadge{
			UserId:  body.UserId,
			BadgeId: body.BadgeId,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create userbadge", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// @Security      BearerAuth
// @Summary 	  Delete UserBadge
// @Description   This Api for deleting userBadge
// @Tags 		  userBadges
// @Accept 		  json
// @Produce 	  json
// @Param 		  UserBadgeDelete body models.UserBadgeDelete true "UserBadgeDelete Model"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/userbadge/{id} [DELETE]
func (h *handlerV1) DeleteUserBadge(ctx *gin.Context) {
	var body models.UserBadgeCreate

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

	response, err := h.storage.UserBadge().Delete(ctxTime, body.UserId, body.BadgeId)
	if err != nil || !response {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotDeletedMessage,
		})
		log.Println("failed to delete user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, true)
}

// // @Security      BearerAuth
// // @Summary       ListUserBadges
// // @Description   This Api for get all userBadges
// // @Tags          userBadges
// // @Accept        json
// // @Produce       json
// // @Param         page query uint64 true "Page"
// // @Param         limit query uint64 true "Limit"
// // @Success 	  200 {object} []models.UserBadgeResponse
// // @Failure		  400 {object} models.Error
// // @Failure		  401 {object} models.Error
// // @Failure		  403 {object} models.Error
// // @Failure       500 {object} models.Error
// // @Router        /v1/userbadges [GET]
// func (h *handlerV1) ListUserBadges(ctx *gin.Context) {
// 	queryParams := ctx.Request.URL.Query()

// 	params, errStr := utils.ParseQueryParams(queryParams)
// 	if errStr != nil {
// 		ctx.JSON(http.StatusInternalServerError, &models.Error{
// 			Message: models.NotFoundMessage,
// 		})
// 		log.Println("failed to parse query params json" + errStr[0])
// 		return
// 	}

// 	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, models.Error{
// 			Message: models.InternalMessage,
// 		})
// 		log.Println("failed to parse timeout", err.Error())
// 		return
// 	}

// 	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
// 	defer cancel()

// 	userBadges, count, err := h.storage.UserBadge().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, &models.Error{
// 			Message: models.NotFoundMessage,
// 		})
// 		log.Println("failed to get all users", err.Error())
// 		return
// 	}
// 	if len(userBadges) == 0 {
// 		ctx.JSON(http.StatusOK, nil)
// 		log.Println("Not found userBadges")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, models.UserBadgeListResponse{
// 		UserBadges: userBadges,
// 		Count:      count,
// 	})
// }

// // @Security      BearerAuth
// // @Summary       ListUserBadges
// // @Description   This Api for get all userBadges
// // @Tags          userBadges
// // @Accept        json
// // @Produce       json
// // @Param         page query uint64 true "Page"
// // @Param         limit query uint64 true "Limit"
// // @Success 	  200 {object} []models.UserBadgeResponse
// // @Failure		  400 {object} models.Error
// // @Failure		  401 {object} models.Error
// // @Failure		  403 {object} models.Error
// // @Failure       500 {object} models.Error
// // @Router        /v1/userbadges [GET]
// func (h *handlerV1) ListUserBadges(ctx *gin.Context) {
// 	queryParams := ctx.Request.URL.Query()

// 	params, errStr := utils.ParseQueryParams(queryParams)
// 	if errStr != nil {
// 		ctx.JSON(http.StatusInternalServerError, &models.Error{
// 			Message: models.NotFoundMessage,
// 		})
// 		log.Println("failed to parse query params json" + errStr[0])
// 		return
// 	}

// 	duration, err := time.ParseDuration(h.cfg.CtxTimeout)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, models.Error{
// 			Message: models.InternalMessage,
// 		})
// 		log.Println("failed to parse timeout", err.Error())
// 		return
// 	}

// 	ctxTime, cancel := context.WithTimeout(context.Background(), duration)
// 	defer cancel()

// 	userBadges, count, err := h.storage.UserBadge().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, &models.Error{
// 			Message: models.NotFoundMessage,
// 		})
// 		log.Println("failed to get all users", err.Error())
// 		return
// 	}
// 	if len(userBadges) == 0 {
// 		ctx.JSON(http.StatusOK, nil)
// 		log.Println("Not found userBadges")
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, models.UserBadgeListResponse{
// 		UserBadges: userBadges,
// 		Count:      count,
// 	})
// }
