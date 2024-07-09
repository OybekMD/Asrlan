package v1

import (
	"net/http"
	"os"
	"path/filepath"

	"log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// File 		upload
// @Security    BearerAuth
// @Summary 	File upload
// @Description File upload
// @Tags 		file-upload
// @Accept 		json
// @Produce 	json
// @Param 		file formData file true "File"
// @Success 	201 {object} string
// @Failure 	400 {object} string
// @Failure 	404 {object} string
// @Router 		/v1/soundupload [post]  // Changed the endpoint to /v1/soundupload
func (h *handlerV1) UploadSoundFile(c *gin.Context) {
	var file File

	err := c.ShouldBind(&file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Println("Error while uploading file", "post", err)
		return
	}

	// Check if the file has a valid sound file extension
	allowedExtensions := []string{".mp3", ".wav", ".ogg"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if filepath.Ext(file.File.Filename) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Couldn't find matching sound file format",
		})
		log.Println("Error while uploading sound file","sound-upload", err)
		return
	}

	id := uuid.New()
	fileName := id.String() + filepath.Ext(file.File.Filename)
	dst, _ := os.Getwd()

	// Update the directory path to "media/sound"
	if _, err := os.Stat(dst + "/media/sound"); os.IsNotExist(err) {
		os.Mkdir(dst+"/media/sound", os.ModePerm)
	}

	filePath := "/media/sound/" + fileName
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

