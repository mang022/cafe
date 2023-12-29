package db

import (
	"database/sql"
	"time"
)

func SelectOwnerByID(id string) (*Owner, error) {
	var o Owner

	if err := CafeDB.QueryRow(
		`
			SELECT *
			FROM owner
			WHERE owner_id = ?
		`,
		id,
	).Scan(&o); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &o, nil
}

func SelectOwnerByPhone(phone string) (*Owner, error) {
	var o Owner

	if err := CafeDB.QueryRow(
		`
			SELECT *
			FROM owner
			WHERE phone = ?
		`,
		phone,
	).Scan(&o); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &o, nil
}

func InsertOwner(o *Owner) error {
	if _, err := CafeDB.Exec(
		`
			INSERT INTO owner (owner_id, phone, salt, password, created_at)
			VALUES (?, ?, ?, ?, ?)
		`,
		o.ID,
		o.Phone,
		o.Salt,
		o.Password,
		time.Now(),
	); err != nil {
		return err
	}

	return nil
}
