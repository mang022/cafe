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
	"github.com/mang022/cafe/dto"
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

	reqBody := dto.SignUpOnwerDto{
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

	reqBody := dto.SignUpOnwerDto{
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

	reqBody := dto.SignUpOnwerDto{
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

	login := dto.SignUpOnwerDto{
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

	reqBody := dto.CreateProductDto{
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

func TestUpdateProduct(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	login := dto.SignUpOnwerDto{
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

	price := 11000
	reqBody := dto.UpdateProductDto{
		Price: &price,
	}
	jsonBody, _ = json.Marshal(reqBody)
	log.Println(string(jsonBody))

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodPut, "/owner/"+claims.OwnerID+"/product/1", bytes.NewBuffer(jsonBody))
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestDeleteProduct(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	login := dto.SignUpOnwerDto{
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

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodDelete, "/owner/"+claims.OwnerID+"/product/2", nil)
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestReadProductDetail(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	login := dto.SignUpOnwerDto{
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

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/owner/"+claims.OwnerID+"/product/1", nil)
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}

func TestReadProductList(t *testing.T) {
	conf.SetupConfig()
	db.SetupDB()
	router := setupRouter()

	login := dto.SignUpOnwerDto{
		Phone:    "010-1234-5679",
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

	w = httptest.NewRecorder()
	req, _ = http.NewRequest(http.MethodGet, "/owner/"+claims.OwnerID+"/product", nil)
	req.Header.Add("Authorization", "Bearer "+jwtToken)
	query := req.URL.Query()
	query.Add("keyword", "ㄱㅊ")
	// query.Add("last_id", "11")
	req.URL.RawQuery = query.Encode()
	router.ServeHTTP(w, req)

	log.Println(w.Body.String())
	assert.Equal(t, http.StatusOK, w.Code)
	db.CloseDB()
}
