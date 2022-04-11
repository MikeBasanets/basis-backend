package algorithm

import (
	"errors"
	"math/rand"
	"time"
)

func makeAnyWardrobe() (Wardrobe, error) {
	wardrobe, err := LoadClothing()
	if err != nil {
		return Wardrobe{}, err
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(wardrobe.Pants), func(i, j int) {
		wardrobe.Pants[i], wardrobe.Pants[j] = wardrobe.Pants[j], wardrobe.Pants[i]
	})
	rand.Shuffle(len(wardrobe.Shirts), func(i, j int) {
		wardrobe.Shirts[i], wardrobe.Shirts[j] = wardrobe.Shirts[j], wardrobe.Shirts[i]
	})
	rand.Shuffle(len(wardrobe.Outerwear), func(i, j int) {
		wardrobe.Outerwear[i], wardrobe.Outerwear[j] = wardrobe.Outerwear[j], wardrobe.Outerwear[i]
	})
	wardrobe.Pants = wardrobe.Pants[0 : 2+rand.Intn(3)]
	wardrobe.Shirts = wardrobe.Shirts[0 : 3+rand.Intn(3)]
	wardrobe.Outerwear = wardrobe.Outerwear[0 : 2+rand.Intn(3)]
	return wardrobe, nil
}

func CalculateWardrobe(q QuizData) (Wardrobe, error) {
	if q.Age >= 0 && q.Age < 255 && len(q.Designation) >= 3 && len(q.Designation) <= 15 {
		return makeAnyWardrobe()
	} else {
		return Wardrobe{}, errors.New("Wrong parameters")
	}
}
