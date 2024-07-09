package v1

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"asrlan-monolight/api/models"

	"github.com/gin-gonic/gin"
)

// @Security      BearerAuth
// @Summary       ListDashboards
// @Description   This Api for get all Dashboard
// @Tags          Dashboard
// @Accept        json
// @Produce       json
// @Param         user_id query string true "UserID"
// @Param         language_id query string true "LanguageId"
// @Param         level_id query string true "LevelId"
// @Success 	  200 {object} models.DashboardResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/dashboard [GET]
func (h *handlerV1) GetDashboard(ctx *gin.Context) {
	fmt.Println("Request  Header:", ctx.GetHeader("Authorization"))
	user_id := ctx.Query("user_id")
	language_id := ctx.Query("language_id")
	level_id := ctx.Query("level_id")

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

	language_id64, err := strconv.Atoi(language_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to get number in lanuage_id:", err.Error())
		return
	}

	level_id64, err := strconv.Atoi(level_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.InternalMessage,
		})
		log.Println("failed to get number in level_id:", err.Error())
		return
	}

	activitys, err := h.storage.Dashboard().GetDashboard(ctxTime, user_id, int64(language_id64), int64(level_id64))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, activitys)
}

// @Security      BearerAuth
// @Summary       ListDashboards
// @Description   This Api for get all Dashboard
// @Tags          Dashboard
// @Accept        json
// @Produce       json
// @Param         user_id query string true "UserID"
// @Success 	  200 {object} models.DashboardResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/navbar [GET]
func (h *handlerV1) GetNavbar(ctx *gin.Context) {
	fmt.Println("Request  Header:", ctx.GetHeader("Authorization"))
	user_id := ctx.Query("user_id")

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

	navbar, err := h.storage.Dashboard().GetNavbar(ctxTime, user_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, navbar)
}

// @Security      BearerAuth
// @Summary       Leaderboard
// @Description   This Api for get all Leaderboard
// @Tags          Leaderboard
// @Accept        json
// @Produce       json
// @Param         period query string true "PeriodSelect"
// @Param         level_id query string true "LevelId"
// @Success 	  200 {object} models.LeaderboardResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/leaderboard [GET]
func (h *handlerV1) GetLeaderboard(ctx *gin.Context) {
	period := ctx.Query("period")
	level_id := ctx.Query("level_id")

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

	activitys, err := h.storage.Dashboard().GetLeaderboard(ctxTime, period, level_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, activitys)
}
