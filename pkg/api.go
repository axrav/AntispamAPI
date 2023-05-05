package pkg

import (
	"net/http"

	"github.com/axrav/antispam/training"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(g *gin.Engine) {
	g.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "online",
		})
	})

	g.POST("/predict", func(c *gin.Context) {
		var data Resp
		c.BindJSON(&data)
		if data.Message == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "bad request",
			})
			return
		}
		scores, isSpam := training.PredictScores(data.Message)
		c.JSON(http.StatusOK, gin.H{
			"message":     data.Message,
			"isSpam":      isSpam,
			"SpamPercent": scores[0],
			"HamPercent":  scores[1],
		})

	})

}

type Resp struct {
	Message string `json:"message"`
}
