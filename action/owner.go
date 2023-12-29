package action

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mang022/cafe/db"
	uuid "github.com/nu7hatch/gouuid"
)

func SignUpOwner(c *gin.Context) {
	var reqBody SignUpOnwerDto
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
