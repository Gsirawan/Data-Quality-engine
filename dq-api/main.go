package main

import (
	"github.com/gin-gonic/gin"
)

type DataRequest struct {
	Data []map[string]interface{} `json:"data"`
}

type ValidationIssue struct {
	Row     int    `json:"row"`
	Field   string `json:"field"`
	Message string `json:"message"`
	Rule    string `json:"rule"`
}

type ValidationReport struct {
	TotalRows   int               `json:"total_rows"`
	ValidRows   int               `json:"valid_rows"`
	InvalidRows int               `json:"invalid_rows"`
	Issues      []ValidationIssue `json:"issues"`
}

func validateData(data []map[string]interface{}) ValidationReport {
	report := ValidationReport{
		TotalRows:   len(data),
		ValidRows:   0,
		InvalidRows: 0,
		Issues:      []ValidationIssue{},
	}
	for i, row := range data {
		rowValid := true
		for column, value := range row {
			if value == nil || value == "" {
				issue := ValidationIssue{
					Row:     i + 1,
					Field:   column,
					Message: "Value is Null or empty",
					Rule:    "Not Null",
				}
				report.Issues = append(report.Issues, issue)
				rowValid = false
			}
		}
		if rowValid {
			report.ValidRows++
		} else {
			report.InvalidRows++
		}

	}
	return report
}

func startServer() {
	router := gin.Default()

	// GET endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "Ok",
		})
	})

	// POST endpoint
	router.POST("/validate", func(c *gin.Context) {
		data := DataRequest{}
		if err := c.BindJSON(&data); err != nil {
			// handle error
			c.JSON(400, gin.H{
				"Error": "Invalid JSON format",
			})

			return

		}
		report := validateData(data.Data)
		c.JSON(200, report)
	})

	// Start server
	router.Run(":8080")
}

// The application!!
func main() {
	startServer()
}
