package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigovilar/go-crud-api/config"
	"github.com/rodrigovilar/go-crud-api/database"
	"github.com/rodrigovilar/go-crud-api/routes"
)

func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run()
}
