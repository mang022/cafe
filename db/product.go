package db

import (
	"database/sql"
	"time"

	"github.com/mang022/cafe/dto"
)

func InsertProduct(p *Product) (int64, error) {
	result, err := CafeDB.Exec(
		`
			INSERT INTO product (owner_id, category, price, cost, name, description, barcode, expiration_time, size, created_at)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		p.OwnerID,
		p.Category,
		p.Price,
		p.Cost,
		p.Name,
		p.Description,
		p.Barcode,
		p.ExpirationTime,
		p.Size,
		time.Now(),
	)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return id, nil
}

func UpdateProduct(pid int64, reqBody dto.UpdateProductDto) error {
	setQuery := ""
	setArg := make([]interface{}, 0)
	if reqBody.Category != nil {
		setQuery += "category = ?, "
		setArg = append(setArg, *reqBody.Category)
	}
	if reqBody.Price != nil {
		setQuery += "price = ?, "
		setArg = append(setArg, *reqBody.Price)
	}
	if reqBody.Cost != nil {
		setQuery += "cost = ?, "
		setArg = append(setArg, *reqBody.Cost)
	}
	if reqBody.Name != nil {
		setQuery += "name = ?, "
		setArg = append(setArg, *reqBody.Name)
	}
	if reqBody.Description != nil {
		setQuery += "description = ?, "
		setArg = append(setArg, *reqBody.Description)
	}
	if reqBody.Barcode != nil {
		setQuery += "barcode = ?, "
		setArg = append(setArg, *reqBody.Barcode)
	}
	if reqBody.ExpirationTime != nil {
		setQuery += "expiration_time = ?, "
		setArg = append(setArg, *reqBody.ExpirationTime)
	}
	if reqBody.Size != nil {
		setQuery += "size = ?, "
		setArg = append(setArg, *reqBody.Size)
	}

	setQuery += "updated_at = ?"
	setArg = append(setArg, time.Now())
	setArg = append(setArg, pid)

	if _, err := CafeDB.Exec(
		`
			UPDATE product
			SET `+setQuery+`
			WHERE product_id = ?
			AND deleted_at IS NULL
		`,
		setArg...,
	); err != nil {
		return err
	}

	return nil
}

func DeleteProductByID(pid int64) error {
	if _, err := CafeDB.Exec(
		`
			UPDATE product
			SET deleted_at = ?
			WHERE product_id = ?
			AND deleted_at IS NULL
		`,
		time.Now(),
		pid,
	); err != nil {
		return err
	}

	return nil
}

func SelectProductByID(id int64) (*Product, error) {
	var p Product
	if err := CafeDB.Get(
		&p,
		`
			SELECT *
			FROM product
			WHERE product_id = ?
			AND deleted_at IS NULL
			LIMIT 1
		`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &p, nil
}
