package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rodrigovilar/go-crud-api/database"
	"github.com/rodrigovilar/go-crud-api/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
