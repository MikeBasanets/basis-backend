package algorithm

import (
	"basis/db"
)

func LoadClothing() (db.Wardrobe, error) {
	pants, err := db.QueryAllPants()
	if err != nil {
		return db.Wardrobe{}, err
	}
	shirts, err := db.QueryAllShirts()
	if err != nil {
		return db.Wardrobe{}, err
	}
	outerwear, err := db.QueryAllOuterwear()
	if err != nil {
		return db.Wardrobe{}, err
	}
	return  db.Wardrobe{Pants: pants, Shirts: shirts, Outerwear: outerwear}, nil
}
