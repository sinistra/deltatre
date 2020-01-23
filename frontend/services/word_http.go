package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sinistra/deltatre/frontend/models"
)

type WordService struct{}

func (w WordService) GetWord(search string) (models.Word, int, error) {

	var word models.Word
	url := fmt.Sprintf("%s/search/%s", os.Getenv("BACKEND_HOST"), search)

	responseBody, statusCode, err := httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(statusCode, err)
		return models.Word{}, statusCode, err
	}

	var apiResponse models.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		log.Println(err)
		return word, http.StatusInternalServerError, err
	}

	// spew.Dump(apiResponse)
	if len(apiResponse.Error) > 0 {
		err = errors.New(apiResponse.Error)
		return models.Word{}, statusCode, err
	}

	word.Text = apiResponse.Data.Text
	word.Count = apiResponse.Data.Count

	return word, statusCode, nil
}

func (w WordService) AddWord(newWord string) (models.Word, int, error) {

	var word models.Word
	url := fmt.Sprintf("%s/add/%s", os.Getenv("BACKEND_HOST"), newWord)

	responseBody, statusCode, err := httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(statusCode, err)
		return models.Word{}, statusCode, err
	}

	var apiResponse models.ApiResponse
	err = json.Unmarshal(responseBody, &apiResponse)
	if err != nil {
		log.Println(err)
		return word, http.StatusInternalServerError, err
	}

	// spew.Dump(apiResponse)
	if len(apiResponse.Error) > 0 {
		err = errors.New(apiResponse.Error)
		return models.Word{}, statusCode, err
	}

	word.Text = apiResponse.Data.Text
	word.Count = apiResponse.Data.Count

	return word, statusCode, nil
}

func (w WordService) TopWords() ([]models.Word, int, error) {
	var words []models.Word

	url := fmt.Sprintf("%s/top", os.Getenv("BACKEND_HOST"))
	responseBody, statusCode, err := httpRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(statusCode, err)
		return words, statusCode, err
	}

	var topResponse models.TopResponse
	err = json.Unmarshal(responseBody, &topResponse)
	if err != nil {
		log.Println(err)
		return words, http.StatusInternalServerError, err
	}

	return topResponse.Data, statusCode, nil
}

func httpRequest(method, url string, body []byte) ([]byte, int, error) {
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
		return nil, http.StatusInternalServerError, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		log.Println(response.StatusCode, err)
		return nil, response.StatusCode, err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(response.StatusCode, err)
		return nil, response.StatusCode, err
	}
	return body, response.StatusCode, nil
}
