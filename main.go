package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fileName := c.PostForm("fileName")

		// Save the uploaded file to a specific location
		dst := fmt.Sprintf("upload/%s", fileName)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Process the file with the provided file name
		// You can call the file processing logic here using 'dst' and 'fileName'

		c.JSON(http.StatusOK, gin.H{"success": true, "message": "File uploaded and processed successfully"})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
