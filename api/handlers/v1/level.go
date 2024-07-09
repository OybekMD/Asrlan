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
// @Summary 	  Create Level
// @Description   This Api for creating a new level
// @Tags 		  levels
// @Accept 		  json
// @Produce 	  json
// @Param 		  LevelCreate body models.LevelCreate true "LevelCreate Model"
// @Success 	  201 {object} models.LevelResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/level [POST]
func (h *handlerV1) CreateLevel(ctx *gin.Context) {
	var body models.LevelCreate

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

	response, err := h.storage.Level().Create(
		ctxTime,
		&repo.Level{
			Name:       body.Name,
			LanguageId: body.LanguageId,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.LevelResponse{
		Id:         response.Id,
		Name:       response.Name,
		LanguageId: response.LanguageId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Level
// @Description   This Api for updating level
// @Tags 		  levels
// @Accept 		  json
// @Produce 	  json
// @Param 		  LevelUpdate body models.LevelUpdate true "Update LevelUpdate Model"
// @Success 	  200 {object} models.LevelResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/level [PUT]
func (h *handlerV1) UpdateLevel(ctx *gin.Context) {
	var body models.LevelUpdate

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

	levelModel := &repo.Level{}
	err = parsing.StructToStruct(&body, levelModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Level().Update(ctxTime, levelModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotUpdatedMessage,
		})
		log.Println("failed to update user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.LevelResponse{
		Id:         response.Id,
		Name:       response.Name,
		LanguageId: response.LanguageId,
		CreatedAt:  response.CreatedAt,
		UpdatedAt:  response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Delete Level
// @Description   This Api for deleting level
// @Tags 		  levels
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/level/{id} [DELETE]
func (h *handlerV1) DeleteLevel(ctx *gin.Context) {
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

	response, err := h.storage.Level().Delete(ctxTime, id)
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
// @Summary 	  Get Level
// @Description   This Api for get level
// @Tags 		  levels
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.LevelResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/level/{id} [GET]
func (h *handlerV1) GetLevel(ctx *gin.Context) {
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

	level, err := h.storage.Level().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get level", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, level)
}

// @Security      BearerAuth
// @Summary       ListLevels
// @Description   This Api for get all levels
// @Tags          levels
// @Accept        json
// @Produce       json
// @Param         page query uint64 true "Page"
// @Param         limit query uint64 true "Limit"
// @Success 	  200 {object} []models.LevelResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/levels [GET]
func (h *handlerV1) ListLevels(ctx *gin.Context) {
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

	levels, count, err := h.storage.Level().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(levels) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found levels")
		return
	}

	ctx.JSON(http.StatusOK, models.LevelListResponse{
		Levels: levels,
		Count:  count,
	})
}

// @Security      BearerAuth
// @Summary       ListLevels
// @Description   This Api for get all levels
// @Tags          levels
// @Accept        json
// @Produce       json
// @Param         language_id query string true "LanguageId"
// @Success 	  200 {object} []models.LevelResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/levelsforregister [GET]
func (h *handlerV1) LevelsForRegister(ctx *gin.Context) {
	language_id := ctx.Query("language_id")

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

	levels, err := h.storage.Level().GetAllForRegister(ctxTime, language_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(levels) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found levels")
		return
	}

	ctx.JSON(http.StatusOK, models.LevelForRegisterResponse{
		Levels: levels,
	})
}

// @Security      BearerAuth
// @Summary       ListLevels
// @Description   This Api for get all levels
// @Tags          levels
// @Accept        json
// @Produce       json
// @Param         user_id query string true "UserId"
// @Param         language_id query string true "LanguageId"
// @Success 	  200 {object} []models.LevelResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/levelsforcourse [GET]
func (h *handlerV1) LevelsForCourse(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	language_id := ctx.Query("language_id")

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

	levels, err := h.storage.Level().GetAllForCourses(ctxTime, user_id, language_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(levels) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found levels")
		return
	}

	ctx.JSON(http.StatusOK, models.LevelForCourseResponse{
		Levels: levels,
	})
}
