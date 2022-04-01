package algorithm

import (
	"basis/db"
)

type QuizData struct {
	Purpose      string `json:"purpose"`
	BirthdayYear int    `json:"birthdayYear"`
}

type Wardrobe struct {
	Pants     []db.Pants     `json:"pants"`
	Shirts    []db.Shirt     `json:"shirts"`
	Outerwear []db.Outerwear `json:"outerwear"`
}
