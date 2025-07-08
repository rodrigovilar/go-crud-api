package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rodrigovilar/go-crud-api/config"
	"github.com/rodrigovilar/go-crud-api/database"
	"github.com/rodrigovilar/go-crud-api/models"
)

func TestMain(m *testing.M) {
	config.LoadEnv()
	code := m.Run()
	os.Exit(code)
}

var testUserID uint

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/api/users", CreateUser)
	r.GET("/api/users", GetUsers)
	r.GET("/api/users/:id", GetUserByID)
	r.PUT("/api/users/:id", UpdateUser)
	r.DELETE("/api/users/:id", DeleteUser)

	database.Connect()
	return r
}

func TestCreateUser(t *testing.T) {
	router := setupTestRouter()

	user := models.User{Name: "Test User", Email: "test@example.com"}
	jsonValue, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/api/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("Expected status 201, got %d", resp.Code)
	}

	var createdUser models.User
	json.Unmarshal(resp.Body.Bytes(), &createdUser)
	testUserID = createdUser.ID
}

func TestGetUsers(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/api/users", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

func TestGetUserByID(t *testing.T) {
	router := setupTestRouter()

	url := fmt.Sprintf("/api/users/%d", testUserID)
	req, _ := http.NewRequest("GET", url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

func TestUpdateUser(t *testing.T) {
	router := setupTestRouter()

	updatedUser := models.User{Name: "Updated User", Email: "updated@example.com"}
	jsonValue, _ := json.Marshal(updatedUser)

	url := fmt.Sprintf("/api/users/%d", testUserID)
	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	router := setupTestRouter()

	url := fmt.Sprintf("/api/users/%d", testUserID)
	req, _ := http.NewRequest("DELETE", url, nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.Code)
	}
}
