package models

type Word struct {
	Text  string `json:text`
	Count int    `json:count`
}

type ApiResponse struct {
	Data    Word   `json:"data"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type TopResponse struct {
	Data    []Word `json:"data"`
	Message string `json:"message"`
}
