package services

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sinistra/deltatre/models"
	"time"
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

func httpRequest(method, url string, body []byte) ([]byte, error) {
	var bodyBuffer io.Reader
	if len(body) > 0 {
		bodyBuffer = bytes.NewBuffer(body)
	}

	client := http.Client{
		Timeout: time.Second * 30,
	}
	request, err := http.NewRequest(method, url, bodyBuffer)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/vnd.api+json")
	request.Header.Set("x-api-key", "ITG.9248ce9dc72946701df7119f02f9ca1e.Xo_Qgyy6ER1uJkAKpCoUo-uvSiXHgD6vtQFh0X1MUmpp35gtdMco7BnEkB3d_726")

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return body, nil
}
