package services

import (
	"errors"
	"fmt"
	"log"
	"sort"

	"sinistra/deltatre/backend/models"
)

var Words = []models.Word{
	{"hello", 0},
	{"goodbye", 0},
	{"simple", 0},
	{"list", 0},
	{"search", 0},
	{"filter", 0},
	{"yes", 0},
	{"no", 0},
}

type WordService struct{}

func (w WordService) GetWords() ([]models.Word, error) {
	return Words, nil
}

func (w WordService) GetWord(search string) (models.Word, error) {

	log.Println(search)
	for i, word := range Words {
		if word.Text == search {
			log.Println("match against ", search)
			Words[i].Count++
			word.Count++
			return word, nil
		}
	}

	err := errors.New("word not found")
	word := models.Word{search, 0}
	return word, err
}

func (w WordService) AddWord(newWord string) (models.Word, error) {
	for _, word := range Words {
		if word.Text == newWord {
			log.Println("word exists ", newWord)
			err := errors.New(fmt.Sprintf("%s is already in the list", newWord))
			return models.Word{}, err
		}
	}

	word := models.Word{newWord, 0}
	Words = append(Words, word)
	return word, nil
}

func (w WordService) TopWords() ([]models.Word, error) {
	// Sort by count, keeping original order or equal elements.
	sort.SliceStable(Words, func(i, j int) bool {
		return Words[i].Count > Words[j].Count
	})
	fmt.Println(Words)

	return Words[:5], nil
}
