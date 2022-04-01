package algorithm

import (
	"basis/db"
)

func LoadClothing() (Wardrobe, error) {
	pants, err := db.QueryAllPants()
	if err != nil {
		return Wardrobe{}, err
	}
	shirts, err := db.QueryAllShirts()
	if err != nil {
		return Wardrobe{}, err
	}
	outerwear, err := db.QueryAllOuterwear()
	if err != nil {
		return Wardrobe{}, err
	}
	return  Wardrobe{Pants: pants, Shirts: shirts, Outerwear: outerwear}, nil
}
