package db

import (
	"database/sql"
)

type Owner struct {
	ID           string        `db:"owner_id"`
	Phone        string        `db:"phone"`
	Salt         string        `db:"salt"`
	Password     string        `db:"password"`
	LastLoginDt  sql.NullInt64 `db:"last_login_dt"`
	LastLogoutDt sql.NullInt64 `db:"last_logout_dt"`
	CreatedAt    []byte        `db:"created_at"`
	UpdatedAt    []byte        `db:"updated_at"`
	DeletedAt    []byte        `db:"deleted_at"`
}
