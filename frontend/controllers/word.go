package controllers

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"sinistra/deltatre/frontend/services"
)

type WordController struct{}

func (w WordController) GetWord(ctx *gin.Context) {
	spew.Dump(ctx.Request.URL.Query())
	// name := ctx.Query("name")
	word := ctx.Param("word")

	service := services.WordService{}

	result, err := service.GetWord(word)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Server error", "error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    result,
	})
}

func (w WordController) AddWord(ctx *gin.Context) {
}

func (w WordController) TopWords(ctx *gin.Context) {
}
