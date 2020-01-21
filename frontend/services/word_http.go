package services

import (
	"../../models"
)

type WordService struct{}

func (w WordService) GetWord(search string) (models.Word, error) {

	var word models.Word

	return word, nil
}

func (w WordService) AddWord(word models.Word) (models.Word, error) {

	return word, nil
}

func (w WordService) TopWords(count int) ([]models.Word, error) {
	var words []models.Word

	return words, nil
}
