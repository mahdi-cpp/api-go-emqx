package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"log"
	"net/http"
)

func AddUploadRoute2(rg *gin.RouterGroup) {

	route := rg.Group("/upload2")

	// Define a POST endpoint to handle file uploads
	route.POST("/", func(c *gin.Context) {

		// Read the uploaded file
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File not found"})
			return
		}

		uniqueID := uuid.New()
		var fileName = "/var/instagram/voices/" + uniqueID.String() + ".mp4"

		// Save the uploaded file to the server
		if err := c.SaveUploadedFile(file, fileName); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
			return
		}

		err = Convert3GPToWAV(fileName, "/var/instagram/voices/"+uniqueID.String()+".wav")
		if err != nil {
			return
		}

		// Read the JSON data from the form
		jsonData := c.PostForm("json")
		if jsonData == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON data not found"})
			return
		}

		var voice repository.Voice
		err = json.Unmarshal([]byte(jsonData), &voice)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}

		voice.FileName = uniqueID.String()

		// Save JSON to file
		err = SaveRecordToFile(voice)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("JSON data saved to data.json")
		}

		// Log the received data
		log.Printf("Received file: %s", file.Filename)
		log.Printf("Received JSON data: %s", jsonData)

		// Respond to the client
		c.JSON(http.StatusOK, gin.H{
			"message":  "File uploaded successfully",
			"filename": file.Filename,
			"json":     jsonData,
		})
	})

}
