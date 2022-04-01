package db

import (
	"context"
)

func QueryAllPants() ([]Pants, error) {
	rows, err := connectionPool.Query(context.Background(), "select id, pageurl, imageurl, color, price from pants")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Pants
	for rows.Next() {
		var i Pants
		if err := rows.Scan(
			&i.ID,
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	/*if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}*/
	return items, nil
}

func QueryAllShirts() ([]Shirt, error) {
	rows, err := connectionPool.Query(context.Background(), "select id, pageurl, imageurl, color, price from shirts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Shirt
	for rows.Next() {
		var i Shirt
		if err := rows.Scan(
			&i.ID,
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}

func QueryAllOuterwear() ([]Outerwear, error) {
	rows, err := connectionPool.Query(context.Background(), "select id, pageurl, imageurl, color, price, warmth from outerwear")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Outerwear
	for rows.Next() {
		var i Outerwear
		if err := rows.Scan(
			&i.ID,
			&i.PageUrl,
			&i.ImageUrl,
			&i.Color,
			&i.Price,
			&i.Warmth,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	return items, nil
}
