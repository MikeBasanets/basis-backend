package clothing

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ClothingItem struct {
	Page   string   `json:"page"`
	Image  string   `json:"image"`
	Style  string   `json:"style"`
	Colors []string `json:"colors"`
	Fit    string   `json:"fit"`
}

type Pants struct {
	ClothingItem
}

type Shirt struct {
	ClothingItem
}

type Outerwear struct {
	ClothingItem
	Warmth int `json:"warmth"`
}

type Wardrobe struct {
	Pants []Pants `json:"pants"`
	Shirts []Shirt `json:"shirts"`
	Outerwear []Outerwear `json:"outerwear"`
}


var AllClothing Wardrobe = Wardrobe{}

func LoadClothing() {
	AllClothing = Wardrobe{Pants: loadPants(), Shirts: loadShirts(), Outerwear: loadOuterwear()}
}

func loadPants() []Pants {
	file, _ := os.Open("data/pants.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var result []Pants
	if json.Unmarshal(byteValue, &result) != nil {
		panic("Cant unmarshal pants")
	}
	return result
}

func loadShirts() []Shirt {
	file, _ := os.Open("data/shirts.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var result []Shirt
	if json.Unmarshal(byteValue, &result) != nil {
		panic("Cant unmarshal shirts")
	}
	return result
}

func loadOuterwear() []Outerwear {
	file, _ := os.Open("data/outerwear.json")
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	var result []Outerwear
	if json.Unmarshal(byteValue, &result) != nil {
		panic("Cant unmarshal outerwear")
	}
	return result
}
