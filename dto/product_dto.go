package dto

type CreateProductDto struct {
	Category       string `form:"category" json:"category" binding:"required,min=1,max=50"`
	Price          int    `form:"price" json:"price" binding:"required,gte=0"`
	Cost           int    `form:"cost" json:"cost" binding:"required,gte=0"`
	Name           string `form:"name" json:"name" binding:"required,min=1,max=200"`
	Description    string `form:"description" json:"description" binding:"required,min=1,max=2000"`
	Barcode        string `form:"barcode" json:"barcode" binding:"required,min=1,max=13"`
	ExpirationTime int    `form:"expiration_time" json:"expiration_time" binding:"required,gt=1"`
	Size           string `form:"size" json:"size" binding:"required,oneof=small large"`
}

type UpdateProductDto struct {
	Category       *string `form:"category" json:"category" binding:"omitempty,min=1,max=50"`
	Price          *int    `form:"price" json:"price" binding:"omitempty,gte=0"`
	Cost           *int    `form:"cost" json:"cost" binding:"omitempty,gte=0"`
	Name           *string `form:"name" json:"name" binding:"omitempty,min=1,max=200"`
	Description    *string `form:"description" json:"description" binding:"omitempty,min=1,max=2000"`
	Barcode        *string `form:"barcode" json:"barcode" binding:"omitempty,min=1,max=13"`
	ExpirationTime *int    `form:"expiration_time" json:"expiration_time" binding:"omitempty,gt=1"`
	Size           *string `form:"size" json:"size" binding:"omitempty,oneof=small large"`
}
