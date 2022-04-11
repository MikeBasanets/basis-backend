package db

import (
	"context"
)

func QueryAllPants() ([]Pants, error) {
	rows, err := connectionPool.Query(context.Background(), `select pageUrl, imageUrl, color, pattern, 
		description, brand, price, season, subcategory, lastUpdated,
		fitType, legOpeningCm from pants`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pants
	for rows.Next() {
		var i Pants
		if err := rows.Scan(
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Pattern,
			&i.Description,
			&i.Brand,
			&i.Price,
			&i.Season,
			&i.Subcategory,
			&i.LastUpdated,
			&i.FitType,
			&i.LegOpeningCm,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func QueryAllShirts() ([]Shirt, error) {
	rows, err := connectionPool.Query(context.Background(), `select pageUrl, imageUrl, color, pattern,
		description, brand, price, season, subcategory, lastUpdated,
		fitType, lengthCm, sleeveLengthCm, collarOrCutout from shirts`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shirt
	for rows.Next() {
		var i Shirt
		if err := rows.Scan(
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Pattern,
			&i.Description,
			&i.Brand,
			&i.Price,
			&i.Season,
			&i.Subcategory,
			&i.LastUpdated,
			&i.FitType,
			&i.LengthCm,
			&i.SleeveLengthCm,
			&i.CollarOrCutout,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func QueryAllOuterwear() ([]Outerwear, error) {
	rows, err := connectionPool.Query(context.Background(), `select pageUrl, imageUrl, color, pattern,
		description, brand, price, season, subcategory, lastUpdated,
		hoodType, lengthCm, sleeveLengthCm, insulationComposition from outerwear`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Outerwear
	for rows.Next() {
		var i Outerwear
		if err := rows.Scan(
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Pattern,
			&i.Description,
			&i.Brand,
			&i.Price,
			&i.Season,
			&i.Subcategory,
			&i.LastUpdated,
			&i.HoodType,
			&i.LengthCm,
			&i.SleeveLengthCm,
			&i.InsulationComposition,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}
