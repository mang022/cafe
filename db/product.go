package db

import (
	"time"
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
