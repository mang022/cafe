package db

import (
	"database/sql"
	"time"
)

type Owner struct {
	ID           string        `db:"owner_id"`
	Phone        string        `db:"phone"`
	Salt         string        `db:"salt"`
	Password     string        `db:"password"`
	LastLoginDt  sql.NullInt64 `db:"last_login_dt"`
	LastLogoutDt sql.NullInt64 `db:"last_logout_dt"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    sql.NullTime  `db:"updated_at"`
	DeletedAt    sql.NullTime  `db:"deleted_at"`
}
