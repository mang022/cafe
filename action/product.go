package action

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mang022/cafe/db"
	"github.com/mang022/cafe/dto"
)

func CreateProduct(c *gin.Context) {
	var reqBody dto.CreateProductDto
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

	id, err := db.InsertProduct(
		&db.Product{
			OwnerID:        c.Param("id"),
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
	var reqBody dto.UpdateProductDto
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

	pid, err := strconv.ParseInt(c.Param("pid"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	if err := db.UpdateProduct(pid, reqBody); err != nil {
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
	})
}

func DeleteProduct(c *gin.Context) {
	pid, err := strconv.ParseInt(c.Param("pid"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	if err := db.DeleteProductByID(pid); err != nil {
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
	})
}

func ReadProductDetail(c *gin.Context) {
	pid, err := strconv.ParseInt(c.Param("pid"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	product, err := db.SelectProductByID(pid)
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
			"product": gin.H{
				"id":              product.ID,
				"category":        product.Category,
				"price":           product.Price,
				"cost":            product.Cost,
				"name":            product.Name,
				"description":     product.Description,
				"barcode":         product.Barcode,
				"expiration_time": product.ExpirationTime,
				"size":            product.Size,
			},
		},
	})
}

func ReadProductList(c *gin.Context) {
	lastID := int64(-1)
	if len(c.Query("last_id")) > 0 {
		var err error
		lastID, err = strconv.ParseInt(c.Query("last_id"), 10, 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"meta": gin.H{
					"code":    http.StatusBadRequest,
					"message": "잘못된 요청입니다.",
				},
			})
			return
		}
	}

	keyword := c.Query("keyword")

	productList, err := db.SelectProductList(c.Param("id"), lastID, keyword)
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

	products := make([]gin.H, 0)
	for _, p := range productList {
		products = append(products, gin.H{
			"id":       p.ID,
			"category": p.Category,
			"price":    p.Price,
			"name":     p.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
		"data": gin.H{
			"products": products,
		},
	})
}
