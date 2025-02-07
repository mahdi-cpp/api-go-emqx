package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

var counter = 0

func AddUploadRoute(rg *gin.RouterGroup) {

	route := rg.Group("/upload")

	// Define the upload endpoint
	route.POST("/", func(c *gin.Context) {

		// Get the file from the request
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
			return
		}

		uniqueID := uuid.New()
		var fileName = "/var/instagram/voices/" + uniqueID.String() + ".mp4"

		// Save the file to the server
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}

		err = Convert3GPToWAV(fileName, "/var/instagram/voices/"+uniqueID.String()+".wav")
		if err != nil {
			return
		}

		// Respond with success message
		//c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %s uploaded successfully", file.Filename)})

		c.String(http.StatusOK, uniqueID.String()+".3gp")
	})

}
