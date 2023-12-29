package action

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mang022/cafe/conf"
	"github.com/mang022/cafe/db"
	"github.com/mang022/cafe/dto"
	uuid "github.com/nu7hatch/gouuid"
)

const ISSUER = "cafe"

type OwnerClaims struct {
	OwnerID string `json:"owner_id"`
	jwt.RegisteredClaims
}

func SignUpOwner(c *gin.Context) {
	var reqBody dto.SignUpOnwerDto
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "휴대폰번호 또는 비밀번호를 다시 확인해주세요.",
			},
		})
		return
	}

	phone := strings.ReplaceAll(reqBody.Phone, "-", "")

	owner, err := db.SelectOwnerByPhone(phone)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"meta": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "나중에 다시 시도해주세요.",
			},
		})
		return
	} else if owner != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "이미 등록된 휴대폰번호입니다.",
			},
		})
		return
	}

	var ownerID *uuid.UUID
	for {
		var err error
		ownerID, err = uuid.NewV4()
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

		owner, err := db.SelectOwnerByID(ownerID.String())
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"meta": gin.H{
					"code":    http.StatusInternalServerError,
					"message": "나중에 다시 시도해주세요.",
				},
			})
			return
		} else if owner == nil {
			break
		}
	}

	salt := genRandomHexStr(16)
	pwd := hashPassword(reqBody.Password, salt)

	if err := db.InsertOwner(
		&db.Owner{
			ID:       ownerID.String(),
			Phone:    phone,
			Salt:     salt,
			Password: pwd,
		},
	); err != nil {
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

func genRandomHexStr(n int) string {
	b := make([]byte, n/2)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func hashPassword(pwd string, salt string) string {
	comb := []byte(pwd + salt)
	sum := sha256.Sum256(comb)
	return hex.EncodeToString(sum[:])
}

func SignInOwner(c *gin.Context) {
	var reqBody dto.SignUpOnwerDto
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "휴대폰번호 또는 비밀번호를 다시 확인해주세요.",
			},
		})
		return
	}

	phone := strings.ReplaceAll(reqBody.Phone, "-", "")

	owner, err := db.SelectOwnerByPhone(phone)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"meta": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "나중에 다시 시도해주세요.",
			},
		})
		return
	} else if owner == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "등록되지 않은 사용자이거나 비밀번호가 일치하지 않습니다.",
			},
		})
		return
	}

	if owner.Password != hashPassword(reqBody.Password, owner.Salt) {
		c.JSON(http.StatusBadRequest, gin.H{
			"meta": gin.H{
				"code":    http.StatusBadRequest,
				"message": "등록되지 않은 사용자이거나 비밀번호가 일치하지 않습니다.",
			},
		})
		return
	}

	if err := db.UpdateOwnerLogin(owner.ID); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"meta": gin.H{
				"code":    http.StatusInternalServerError,
				"message": "나중에 다시 시도해주세요.",
			},
		})
		return
	}

	token, err := genJwtToken(owner.ID)

	c.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "ok",
		},
		"data": gin.H{
			"jwt": token,
		},
	})
}

func genJwtToken(ownerID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["owner_id"] = ownerID
	claims["iss"] = ISSUER
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(conf.Conf.JWT.Secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func SignOutOwner(c *gin.Context) {
	ownerID := c.Param("id")

	if err := db.UpdateOwnerLogout(ownerID); err != nil {
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
