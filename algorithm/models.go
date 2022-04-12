package algorithm

import (
	"basis/db"
)

type QuizData struct {
	Designation          string
	Age                  int
	Budget               int
	HairColor            string
	FavoriteColorScheme string
	PreferredFit         string
}

type Wardrobe struct {
	Pants     []db.Pants     `json:"pants"`
	Shirts    []db.Shirt     `json:"shirts"`
	Outerwear []db.Outerwear `json:"outerwear"`
}
