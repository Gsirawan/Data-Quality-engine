package main

import (
	"github.com/gin-gonic/gin"
)

type DataRequest struct {
	Data []map[string]interface{} `json:"data"`
}

func startServer() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Ok",
		})
	})

	// Start server
	router.Run(":8080")
}

// The application!!
func main() {
	startServer()
}
