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
