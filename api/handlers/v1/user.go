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
// @Summary 	  Update User
// @Description   This Api for updating user
// @Tags 		  users
// @Accept 		  json
// @Produce 	  json
// @Param 		  User body models.UserUpdate true "Update User Model"
// @Success 	  200 {object} models.UserResponse
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/user [PUT]
func (h *handlerV1) UpdateUser(ctx *gin.Context) {
	var body models.UserUpdate

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

	// Validate start
	err = body.ValidateEmpity()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Error{
			Message: models.WrongInfoMessage,
		})
		log.Println("Error validating user. id or username not given", err.Error())
		return
	}
	// Validate end

	// Noexistense Start

	responseUsername, err := h.storage.User().CheckUsername(
		ctxTime, body.Id, body.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while getting exist fild",
		})
		log.Println("Error while getting exist fild: ", err.Error())
		return
	}
	if !responseUsername {
		// 409 it means username already exist
		ctx.JSON(http.StatusConflict, gin.H{
			"error": "The username already exists. Please choose a different username.",
		})
		log.Println("Username already use:", body.Username)
		return
	}
	// Noexistense End

	userModel := &repo.User{}
	err = parsing.StructToStruct(&body, userModel)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.InternalMessage,
		})
		log.Println("Error parsing struct to struct", err.Error())
		return
	}

	response, err := h.storage.User().Update(ctxTime, userModel)
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
// @Summary 	  Update User
// @Description   This Api for updating user
// @Tags 		  users
// @Accept 		  json
// @Produce 	  json
// @Param 		  User body models.UserUpdatePassword true "Update User Model"
// @Success 	  200 {object} bool
// @Failure 	  400 {object} models.Error
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/userpassword [PUT]
func (h *handlerV1) UpdateUserPassword(ctx *gin.Context) {
	var body models.UserUpdatePassword

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

	// Noexistense Start

	responseUsername, err := h.storage.User().UpdatePassword(
		ctxTime, &repo.UserUpdatePassword{
			Id:          body.Id,
			OldPassword: body.OldPassword,
			NewPassword: body.NewPassword,
		})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error while getting exist fild",
		})
		log.Println("Error while getting exist fild: ", err.Error())
		return
	}
	if responseUsername == false {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect password or id",
		})
		log.Println("Incorrect password or id", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, responseUsername)
}

// @Security      BearerAuth
// @Summary 	  Delete User
// @Description   This Api for deleting user
// @Tags 		  users
// @Accept 		  json
// @Produce 	  json
// @Param         id path string true "ID"
// @Success 	  200 {object} bool
// @Failure 	  401 {object} models.Error
// @Failure 	  403 {object} models.Error
// @Failure 	  500 {object} models.Error
// @Router 		  /v1/user/{id} [DELETE]
func (h *handlerV1) DeleteUser(ctx *gin.Context) {
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

	response, err := h.storage.User().Delete(ctxTime, id)
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
// @Summary 	  Get User
// @Description   This Api for get user
// @Tags 		  users
// @Accept        json
// @Produce       json
// @Param         id path string true "ID"
// @Success 	  200 {object} models.UserResponse
// @Failure		  401 {object} models.Error
// @Failure		  403 {object} models.Error
// @Failure       500 {object} models.Error
// @Router        /v1/user/{id} [GET]
func (h *handlerV1) GetUser(ctx *gin.Context) {
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

	user, err := h.storage.User().Get(ctxTime, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Error{
			Message: models.NotFoundMessage,
		})
		log.Println("failed to get user", err.Error())
		return
	}

	ctx.JSON(http.StatusOK, user)
}
