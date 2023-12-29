package action

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mang022/cafe/db"
)

func CreateProduct(c *gin.Context) {
	var reqBody CreateProductDto
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "상품 정보를 다시 확인해주세요.",
			},
		})
		return
	}

	ownerID, ok := c.Get("owner_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	id, err := db.InsertProduct(
		&db.Product{
			OwnerID:        ownerID.(string),
			Category:       reqBody.Category,
			Price:          reqBody.Price,
			Cost:           reqBody.Cost,
			Name:           reqBody.Name,
			Description:    reqBody.Description,
			Barcode:        reqBody.Barcode,
			ExpirationTime: reqBody.ExpirationTime,
			Size:           reqBody.Size,
		},
	)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"meta": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "나중에 다시 시도해주세요.",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
		"data": gin.H{
			"id": id,
		},
	})
}

func UpdateProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
	})
}

func DeleteProduct(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
	})
}

func ReadProductList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
	})
}

func ReadProductDetail(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
	})
}
