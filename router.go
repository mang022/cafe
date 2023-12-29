package main

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mang022/cafe/action"
	"github.com/mang022/cafe/conf"
	"github.com/mang022/cafe/dto"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", dto.ValidatePhone)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", action.SignUpOwner)
	r.POST("/signin", action.SignInOwner)

	ownerR := r.Group("/owner/:id")
	ownerR.Use(checkJwtToken)
	ownerR.Use(checkOwner).POST("/signout", action.SignOutOwner)

	ownerR.Use(checkOwner).POST("/product", action.CreateProduct)
	ownerR.Use(checkOwner).PUT("/product/:pid", action.UpdateProduct)
	ownerR.Use(checkOwner).DELETE("/product/:pid", action.DeleteProduct)
	ownerR.Use(checkOwner).GET("/product", action.ReadProductList)
	ownerR.Use(checkOwner).GET("/product/:pid", action.ReadProductDetail)

	return r
}

func checkJwtToken(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	expiredTime, err := claims.GetExpirationTime()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	if expiredTime.Unix() < time.Now().Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	c.Next()
}

func extractBearerToken(header string) (string, error) {
	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.New("incorrectly formatted authorization header")
	}

	return jwtToken[1], nil
}

func parseToken(jwtToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(
		jwtToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid sign")
			}

			return []byte(conf.Conf.JWT.Secret), nil
		},
	)
	if err != nil {
		return nil, errors.New("bad jwt token")
	}

	return token, nil
}

func checkOwner(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader("Authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	token, err := parseToken(jwtToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	signedID, ok := claims["owner_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	reqID := c.Param("id")
	if signedID != reqID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"meta": gin.H{
				"code":    http.StatusUnauthorized,
				"message": "잘못된 요청입니다.",
			},
		})
		return
	}

	c.Next()
}
