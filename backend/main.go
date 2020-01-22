package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"sinistra/deltatre/backend/controllers"
)

func main() {

	wordController := controllers.WordController{}

	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// custom format
		return fmt.Sprintf("%s |%s|[%s]|%s|%s|%d|%s|%s|%s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Method,
			param.ClientIP,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		v1.GET("/words", wordController.GetWords)
		v1.GET("/search/:word", wordController.GetWord)
		v1.GET("/add/:word", wordController.AddWord)
		v1.GET("/top", wordController.TopWords)
	}

	// healthcheck
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Listen and serve on 0.0.0.0:8001
	router.Run(":8001")
}
