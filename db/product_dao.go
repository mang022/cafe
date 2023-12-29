package db

type Product struct {
	ID             int64  `db:"product_id"`
	OwnerID        string `db:"owner_id"`
	Category       string `db:"category"`
	Price          int    `db:"price"`
	Cost           int    `db:"cost"`
	Name           string `db:"name"`
	Description    string `db:"description"`
	Barcode        string `db:"barcode"`
	ExpirationTime int    `db:"expiration_time"`
	Size           string `db:"size"`
	CreatedAt      []byte `db:"created_at"`
	UpdatedAt      []byte `db:"updated_at"`
	DeletedAt      []byte `db:"deleted_at"`
}
