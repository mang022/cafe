package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mang022/cafe/action"
	"github.com/mang022/cafe/conf"
	"github.com/mang022/cafe/db"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestSignup(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	reqBody := action.SignUpOnwerDto{
		Phone:    "010-1234-5678",
		Password: "12345678",
	}
	jsonBody, _ := json.Marshal(reqBody)
	log.Println(string(jsonBody))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestSignin(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	reqBody := action.SignUpOnwerDto{
		Phone:    "010-1234-5678",
		Password: "12345678",
	}
	jsonBody, _ := json.Marshal(reqBody)
	log.Println(string(jsonBody))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestSignout(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	reqBody := action.SignUpOnwerDto{
		Phone:    "010-1234-5678",
		Password: "12345678",
	}
	jsonBody, _ := json.Marshal(reqBody)
	log.Println(string(jsonBody))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	resp := make(map[string]interface{})
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	data := resp["data"].(map[string]interface{})
	jwtToken := data["jwt"].(string)

	token, err := jwt.ParseWithClaims(jwtToken, &action.OwnerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.JWT.Secret), nil
	})
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(*action.OwnerClaims)
	if !ok {
		panic(errors.New("invalid claims"))
	}

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/owner/"+claims.OwnerID+"/signout", nil)
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestCreateProduct(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	login := action.SignUpOnwerDto{
		Phone:    "010-1234-5678",
		Password: "12345678",
	}
	jsonBody, _ := json.Marshal(login)
	log.Println(string(jsonBody))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/signin", bytes.NewBuffer(jsonBody))
	router.ServeHTTP(w, req)

	resp := make(map[string]interface{})
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		panic(err)
	}

	data := resp["data"].(map[string]interface{})
	jwtToken := data["jwt"].(string)

	token, err := jwt.ParseWithClaims(jwtToken, &action.OwnerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.JWT.Secret), nil
	})
	if err != nil {
		panic(err)
	}

	claims, ok := token.Claims.(*action.OwnerClaims)
	if !ok {
		panic(errors.New("invalid claims"))
	}

	reqBody := action.CreateProductDto{
		Category:       "음식",
		Price:          10000,
		Cost:           4000,
		Name:           "김치찜",
		Description:    "100년 전통의 김치찜",
		Barcode:        "1234567890123",
		ExpirationTime: 24 * 7,
		Size:           "small",
	}
	jsonBody, _ = json.Marshal(reqBody)
	log.Println(string(jsonBody))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPost, "/owner/"+claims.OwnerID+"/product", bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}
