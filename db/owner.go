package db

import (
	"database/sql"
	"time"
)

func SelectOwnerByID(id string) (*Owner, error) {
	var o Owner
	if err := CafeDB.Get(
		&o,
		`
			SELECT *
			FROM owner
			WHERE owner_id LIKE ?
			LIMIT 1
		`,
		id,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &o, nil
}

func SelectOwnerByPhone(phone string) (*Owner, error) {
	var o Owner
	if err := CafeDB.Get(
		&o,
		`
			SELECT *
			FROM owner
			WHERE phone LIKE ?
			LIMIT 1
		`,
		phone,
	); err != nil {
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

func UpdateOwnerLogin(id string) error {
	if _, err := CafeDB.Exec(
		`
			UPDATE owner
			SET last_login_dt = ?, updated_at = ?
			WHERE owner_id LIKE ?
		`,
		time.Now().Unix(),
		time.Now(),
		id,
	); err != nil {
		return err
	}

	return nil
}

func UpdateOwnerLogout(id string) error {
	if _, err := CafeDB.Exec(
		`
			UPDATE owner
			SET last_logout_dt = ?, updated_at = ?
			WHERE owner_id LIKE ?
		`,
		time.Now().Unix(),
		time.Now(),
		id,
	); err != nil {
		return err
	}

	return nil
}
