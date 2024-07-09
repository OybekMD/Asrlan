package v1

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

type File struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}


// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Param  img_name query string true "File Name"
// @Router /v1/badgeupload [post]  // Changed the endpoint to /v1/badgeupload
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
func (h *handlerV1) BadgeImageFile(c *gin.Context) {
	var file File
	img_name := c.Query("img_name")

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("Error while uploading file", "post", err)
		return
	}

	// Check if the file has a valid image file extension
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching image file format",
		})
		log.Println("Error while uploading image file", "image-upload", err)
		return
	}

	fileName := img_name + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	// Update the directory path to "/media/images/badges"
	if _, err := os.Stat(dst + "/media/images/badges"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/images/badges", os.ModePerm)
	}

	filePath := "/media/images/badges/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couldn't find matching information, Have you registered before?",
		})
		log.Println("Error while getting customer by email", "post", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}


// File upload
// @Security    BearerAuth
// @Summary File upload
// @Description File upload
// @Tags file-upload
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Router /v1/avatarupload [post]  // Changed the endpoint to /v1/avatarupload
// @Success 201 {object} string
// @Failure 400 {object} string
// @Failure 404 {object} string
func (h *handlerV1) AvatarImageFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("Error while uploading file", "post", err)
		return
	}

	// Check if the file has a valid image file extension
	allowedExtensions := []string{".png", ".jpg", ".jpeg"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching image file format",
		})
		log.Println("Error while uploading image file", "image-upload", err)
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	// Update the directory path to "/media/images/avatars"
	if _, err := os.Stat(dst + "/media/images/avatars"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/images/avatars", os.ModePerm)
	}

	filePath := "/media/images/avatars/" + fileName
	err = c.SaveUploadedFile(file.File, dst+filePath)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Couldn't find matching information, Have you registered before?",
		})
		log.Println("Error while getting customer by email", "post", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"url": c.Request.Host + filePath,
	})
}
