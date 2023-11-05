package main

import (
	"chatgpt-api/pkg/chatgpt"
	"chatgpt-api/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	gptClient, err := chatgpt.Init()
	if err != nil {
		logger.Log(logger.Error, err.Error())
		os.Exit(1)
	}

	// ROUTES
	router.POST("/askGPT", func(c *gin.Context) {
		if gptClient.CheckAPIKey(c.GetHeader("APIKey")) == false {
			logger.Log(logger.Error, "invalid APIKey provided")
			c.String(http.StatusUnauthorized, "invalid API Key provided")
		} else {
			query, _ := c.GetPostForm("question")

			res, err := gptClient.AskGPT(query)
			if err != nil {
				logger.Log(logger.Error, err.Error())
			}

			c.String(http.StatusOK, *res)
		}
	})

	// Run the API on port 8080
	err = router.Run(":8080")
	if err != nil {
		return
	}
}
