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
// @Summary 	  Create Language
// @Description   This Api for creating a new language
// @Tags 		  languages
// @Accept 		  json
// @Produce 	  json
// @Param 		  LanguageCreate body models.LanguageCreate true "LanguageCreate Model"
// @Success 	  201 {object} models.LanguageResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/language [POST]
func (h *handlerV1) CreateLanguage(ctx *gin.Context) {
	var body models.LanguageCreate

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

	response, err := h.storage.Language().Create(
		ctxTime,
		&repo.Language{
			Name:    body.Name,
			Picture: body.Picture,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotCreatedMessage,
		})
		log.Println("failed to create user", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &models.LanguageResponse{
		Id:           response.Id,
		Name:         response.Name,
		Picture: response.Picture,
		CreatedAt: response.CreatedAt,
		UpdatedAt:      response.UpdatedAt,
	})
}

// @Security      BearerAuth
// @Summary 	  Update Language
// @Description   This Api for updating language
// @Tags 		  languages
// @Accept 		  json
// @Produce 	  json
// @Param 		  LanguageUpdate body models.LanguageUpdate true "Update LanguageUpdate Model"
// @Success 	  200 {object} models.LanguageResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/language [PUT]
func (h *handlerV1) UpdateLanguage(ctx *gin.Context) {
	var body models.LanguageUpdate

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

	languageModel := &repo.Language{}
	err = parsing.StructToStruct(&body, languageModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.Language().Update(ctxTime, languageModel)
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
// @Summary 	  Delete Language
// @Description   This Api for deleting language
// @Tags 		  languages
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/language/{id} [DELETE]
func (h *handlerV1) DeleteLanguage(ctx *gin.Context) {
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

	response, err := h.storage.Language().Delete(ctxTime, id)
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
// @Summary 	  Get Language
// @Description   This Api for get language
// @Tags 		  languages
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.LanguageResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/language/{id} [GET]
func (h *handlerV1) GetLanguage(ctx *gin.Context) {
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

	language, err := h.storage.Language().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get language", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, language)
}

// @Security      BearerAuth
// @Summary       ListLanguages
// @Description   This Api for get all languages
// @Tags          languages
// @Accept        json
// @Produce       json
// @Param         page query uint64 true "Page"
// @Param         limit query uint64 true "Limit"
// @Success 	  200 {object} []models.LanguageResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/languages [GET]
func (h *handlerV1) ListLanguages(ctx *gin.Context) {
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

	languages, count, err := h.storage.Language().GetAll(ctxTime, uint64(params.Page), uint64(params.Limit))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(languages) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found languages")
		return
	}

	ctx.JSON(http.StatusOK, models.LanguageListResponse{
		Languages: languages,
		Count:     count,
	})
}

// @Security      BearerAuth
// @Summary       ListLanguages
// @Description   This Api for get all languages
// @Tags          languages
// @Accept        json
// @Produce       json
// @Success 	  200 {object} models.LanguageForRegisterResponse
// @Failure		  400 {object} models.Error
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/languagesforregister [GET]
func (h *handlerV1) LanguagesForRegister(ctx *gin.Context) {
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

	languages, err := h.storage.Language().GetAllForRegister(ctxTime)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get all users", err.Error())
		return
	}
	if len(languages) == 0 {
		ctx.JSON(http.StatusOK, nil)
		log.Println("Not found languages")
		return
	}

	ctx.JSON(http.StatusOK, models.LanguageForRegisterResponse{
		Languages: languages,
	})
}
