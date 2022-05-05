package algorithm

import (
	"basis/db"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func CalculateWardrobe(q db.QuizData) (db.Wardrobe, error) {
	if !(q.Age >= 0 && q.Age < 200 && len(q.Designation) >= 3 && len(q.Designation) <= 15) {
		return db.Wardrobe{}, errors.New("Wrong parameters")
	}
	clothing, err := LoadClothing()
	if err != nil {
		return db.Wardrobe{}, err
	}
	budgetLeft := q.Budget
	if budgetLeft < 1000 {
		budgetLeft = 1000
	}
	clothing = filterCost(clothing, budgetLeft)
	fmt.Println(len(clothing.Outerwear))
	fmt.Println(len(clothing.Pants))
	fmt.Println(len(clothing.Shirts))
	clothing = filterRandomSubset(clothing)
	return clothing, nil
}

func filterCost(w db.Wardrobe, budget int) db.Wardrobe {
	result := db.Wardrobe{}
	for _, v := range w.Outerwear {
		if v.Price <= budget/5 {
			result.Outerwear = append(result.Outerwear, v)
		}
	}
	for _, v := range w.Pants {
		if v.Price <= budget/10 {
			result.Pants = append(result.Pants, v)
		}
	}
	for _, v := range w.Shirts {
		if v.Price <= budget/14 {
			result.Shirts = append(result.Shirts, v)
		}
	}
	return result
}

func filterRandomSubset(w db.Wardrobe) db.Wardrobe {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(w.Pants), func(i, j int) {
		w.Pants[i], w.Pants[j] = w.Pants[j], w.Pants[i]
	})
	rand.Shuffle(len(w.Shirts), func(i, j int) {
		w.Shirts[i], w.Shirts[j] = w.Shirts[j], w.Shirts[i]
	})
	rand.Shuffle(len(w.Outerwear), func(i, j int) {
		w.Outerwear[i], w.Outerwear[j] = w.Outerwear[j], w.Outerwear[i]
	})
	w.Pants = w.Pants[0 : 2+rand.Intn(3)]
	w.Shirts = w.Shirts[0 : 3+rand.Intn(3)]
	w.Outerwear = w.Outerwear[0 : 2+rand.Intn(3)]
	return w
}
