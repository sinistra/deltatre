package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/sinistra/deltatre/backend/services"
)

type WordController struct{}

func (w WordController) GetWord(ctx *gin.Context) {
	word := strings.ToLower(ctx.Param("word"))

	service := services.WordService{}
	result, err := service.GetWord(word)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}
func (w WordController) GetWords(ctx *gin.Context) {

	service := services.WordService{}
	result, err := service.GetWords()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}

func (w WordController) AddWord(ctx *gin.Context) {
	word := strings.ToLower(ctx.Param("word"))

	service := services.WordService{}
	result, err := service.AddWord(word)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Cannot add word", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}

func (w WordController) TopWords(ctx *gin.Context) {
	service := services.WordService{}
	result, err := service.TopWords()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})

}
