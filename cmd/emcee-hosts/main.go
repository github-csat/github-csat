package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/submit", func(c *gin.Context) {
		issueURL := c.Query("issue")
		responseValue := c.Query("response")

		c.String(http.StatusOK, "Your response on %s has been recorded as %s", issueURL, responseValue)
	})
	router.Run(":8080")
}
