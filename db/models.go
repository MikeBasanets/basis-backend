package db

import "time"

type ClothingItem struct {
	PageUrl     string    `json:"pageUrl"`
	ImageUrl    string    `json:"imageUrl"`
	Color       string    `json:"-"`
	Pattern     string    `json:"-"`
	Description string    `json:"description"`
	Brand       string    `json:"brand"`
	Price       int       `json:"price"`
	Season      string    `json:"-"`
	Subcategory string    `json:"-"`
	LastUpdated time.Time `json:"-"`
}

type Outerwear struct {
	ClothingItem
	HoodType              string `json:"-"`
	LengthCm              int    `json:"-"`
	SleeveLengthCm        int    `json:"-"`
	InsulationComposition string `json:"-"`
}

type Pants struct {
	ClothingItem
	FitType      string `json:"-"`
	LegOpeningCm int    `json:"-"`
}

type Shirt struct {
	ClothingItem
	FitType        string `json:"-"`
	LengthCm       int    `json:"-"`
	SleeveLengthCm int    `json:"-"`
	CollarOrCutout string `json:"-"`
}

type User struct {
	Id           int64
	Username     string
	PasswordHash string
	LastActive   time.Time
}

type QuizData struct {
	Designation         string
	Age                 int
	Budget              int
	HairColor           string
	FavoriteColorScheme string
	PreferredFit        string
}

type Wardrobe struct {
	Pants     []Pants     `json:"pants"`
	Shirts    []Shirt     `json:"shirts"`
	Outerwear []Outerwear `json:"outerwear"`
}
