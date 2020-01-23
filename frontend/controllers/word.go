package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/sinistra/deltatre/frontend/services"
)

type WordController struct{}

func (w WordController) GetWord(ctx *gin.Context) {
	word := ctx.Param("word")

	service := services.WordService{}
	result, statusCode, err := service.GetWord(word)

	if err != nil {
		ctx.JSON(statusCode, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"message": "ok",
		"data":    result,
	})
}

func (w WordController) AddWord(ctx *gin.Context) {
	word := ctx.Param("word")

	service := services.WordService{}
	result, statusCode, err := service.AddWord(word)

	if err != nil {
		ctx.JSON(statusCode, gin.H{"status": statusCode, "error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"message": "ok",
		"data":    result,
		"status":  statusCode,
	})

}

func (w WordController) TopWords(ctx *gin.Context) {
	service := services.WordService{}
	result, statusCode, err := service.TopWords()
	if err != nil {
		ctx.JSON(statusCode, gin.H{"status": statusCode, "error": err.Error()})
		return
	}

	ctx.JSON(statusCode, gin.H{
		"message": "ok",
		"data":    result,
		"status":  statusCode,
	})
}
