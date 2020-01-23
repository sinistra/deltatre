package models

type Word struct {
	Text  string `json:text`
	Count int    `json:count`
}
