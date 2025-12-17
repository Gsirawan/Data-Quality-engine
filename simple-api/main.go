package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	displayMenu = "(1). JSON Call \n(2). String Call \n(3). HTML Call \n(4). XML call \n(5). Data Call \n(6). File Call \n(7). Redirect Call \n(8). Exit"
	filePath    = "contacts.txt"
)

func json() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Me": "Hello JSON",
		})
	})

	// Start server
	router.Run(":8080")
}

func string() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "Hello String")
	})

	// Start server
	router.Run(":8080")
}

func html() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("<h1>Hello HTML</h1>"))
	})

	// Start server
	router.Run(":8080")
}

func xml() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.XML(200, gin.H{
			"Me":    "Hello World",
			"World": "Go away",
		})
	})

	// Start server
	router.Run(":8080")
}

func data() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.Data(200, "text/html; charset=utf-8", []byte("<H1>Hello Data</H1>"))
	})

	// Start server
	router.Run(":8080")
}

func file() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.File(filePath)
	})

	// Start server
	router.Run(":8080")
}

func redirect() {
	// Create router
	router := gin.Default()

	// Simple health check endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.Redirect(302, "hhtps://www.google.com")
	})

	// Start server
	router.Run(":8080")
}

// The application!!
func main() {
	choice := "0"
	for {
		fmt.Println(displayMenu)
		fmt.Println("=====================")
		fmt.Println("Please choice from the list:")
		fmt.Println("---------------------")
		_, err := fmt.Scanf("%s", &choice)
		if err != nil {
			log.Fatalf("Error reading input %v\n:", err)
			return
		}
		if choice == "1" {
			json()
		} else if choice == "2" {
			string()
		} else if choice == "3" {
			html()
		} else if choice == "4" {
			xml()
		} else if choice == "5" {
			data()
		} else if choice == "6" {
			file()
		} else if choice == "7" {
			redirect()
		} else if choice == "8" {
			fmt.Println("Goodbye!")
			break
		} else {
			fmt.Println("You can only choose from the List")
			// fmt.Scanln(&choice)
		}
	}
}
