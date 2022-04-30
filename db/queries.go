package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
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

func QueryUserByUsername(username string) (User, error) {
	row := connectionPool.QueryRow(
		context.Background(),
		"select id, username, passwordHash, lastActive from users where username=$1",
		username)
	var user User
	if err := row.Scan(
		&user.Id,
		&user.Username,
		&user.PasswordHash,
		&user.LastActive,
	); err != nil {
		return User{}, err
	}
	return user, nil
}

func SaveUser(user User) error {
	_, err := connectionPool.Exec(context.Background(),
		`
			INSERT INTO users (username, passwordHash, lastActive)
			    VALUES ($1, $2, $3)`,
		user.Username,
		user.PasswordHash,
		user.LastActive)
	return err
}

func SaveResult(result Wardrobe, username string) error {
	transaction, err := connectionPool.BeginTx(context.TODO(), pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer transaction.Rollback(context.TODO())
	row := transaction.QueryRow(context.Background(),
		"select id from users where username = $1",
		username)
	var userId int64
	row.Scan(&userId)
	var resultId int64
	err = transaction.QueryRow(context.Background(),
		"insert into results (userId, resultTime) values ($1, $2) returning id",
		userId, time.Now()).Scan(&resultId)
	if err != nil {
		return err
	}
	err = transaction.Commit(context.TODO())
	for _, v := range result.Outerwear {
		transaction.Exec(context.Background(),
			`
				INSERT INTO resultOuterwear (itemId, resultId)
				    VALUES ($1, $2)`,
			v.PageUrl,
			resultId)
	}
	for _, v := range result.Shirts {
		transaction.Exec(context.Background(),
			`
				INSERT INTO resultShirts (itemId, resultId)
				    VALUES ($1, $2)`,
			v.PageUrl,
			resultId)
	}
	for _, v := range result.Pants {
		transaction.Exec(context.Background(),
			`
				INSERT INTO resultPants (itemId, resultId)
				    VALUES ($1, $2)`,
			v.PageUrl,
			resultId)
	}
	return nil
}

func QueryResultsByDateByUsername(username string) (map[time.Time]*Wardrobe, error) {
	resultsByDate := make(map[time.Time]*Wardrobe)
	//
	rows, err := connectionPool.Query(context.Background(), `
	select results.resultTime, pageUrl, imageUrl, color, pattern, description, brand, price, season, subcategory, lastUpdated, fitType, lengthCm, sleeveLengthCm, collarOrCutout from shirts join resultShirts on shirts.pageUrl = resultShirts.itemId join results on resultShirts.resultId = results.id join users on
results.userid = users.id where username= $1`, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t time.Time
		var s Shirt
		if err := rows.Scan(
			&t,
			&s.PageUrl,
			&s.ImageUrl,
			&s.Color,
			&s.Pattern,
			&s.Description,
			&s.Brand,
			&s.Price,
			&s.Season,
			&s.Subcategory,
			&s.LastUpdated,
			&s.FitType,
			&s.LengthCm,
			&s.SleeveLengthCm,
			&s.CollarOrCutout,
		); err != nil {
			return nil, err
		}
		_, present := resultsByDate[t]
		if !present {
			resultsByDate[t] = &Wardrobe{}
		}
		wardrobe := resultsByDate[t]
		wardrobe.Shirts = append(wardrobe.Shirts, s)
	}
	rows.Close()
	//
	rows, err = connectionPool.Query(context.Background(), `
		select results.resultTime, pageUrl, imageUrl, color, pattern, 
		description, brand, price, season, subcategory, lastUpdated,
		fitType, legOpeningCm from pants join resultPants on pants.pageUrl = resultPants.itemId join results on resultPants.resultId = results.id join users on
		results.userid = users.id where username= $1`, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t time.Time
		var p Pants
		if err := rows.Scan(
			&t,
			&p.PageUrl,
			&p.ImageUrl,
			&p.Color,
			&p.Pattern,
			&p.Description,
			&p.Brand,
			&p.Price,
			&p.Season,
			&p.Subcategory,
			&p.LastUpdated,
			&p.FitType,
			&p.LegOpeningCm,
		); err != nil {
			return nil, err
		}
		_, present := resultsByDate[t]
		if !present {
			resultsByDate[t] = &Wardrobe{}
		}
		wardrobe := resultsByDate[t]
		wardrobe.Pants = append(wardrobe.Pants, p)
	}
	rows.Close()
	//
	rows, err = connectionPool.Query(context.Background(), `
		select results.resultTime, pageUrl, imageUrl, color, pattern,
		description, brand, price, season, subcategory, lastUpdated,
		hoodType, lengthCm, sleeveLengthCm, insulationComposition from outerwear join resultOuterwear on outerwear.			pageUrl = resultOuterwear.itemId join results on resultOuterwear.resultId = results.id join users on
		results.userid = users.id where username= $1`, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t time.Time
		var o Outerwear
		if err := rows.Scan(
			&t,
			&o.PageUrl,
			&o.ImageUrl,
			&o.Color,
			&o.Pattern,
			&o.Description,
			&o.Brand,
			&o.Price,
			&o.Season,
			&o.Subcategory,
			&o.LastUpdated,
			&o.HoodType,
			&o.LengthCm,
			&o.SleeveLengthCm,
			&o.InsulationComposition,
		); err != nil {
			return nil, err
		}
		_, present := resultsByDate[t]
		if !present {
			resultsByDate[t] = &Wardrobe{}
		}
		wardrobe := resultsByDate[t]
		wardrobe.Outerwear = append(wardrobe.Outerwear, o)
	}
	rows.Close()
	//
	return resultsByDate, nil
}
