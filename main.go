package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigovilar/go-crud-api/config"
	"github.com/rodrigovilar/go-crud-api/database"
	_ "github.com/rodrigovilar/go-crud-api/docs"
	"github.com/rodrigovilar/go-crud-api/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Go CRUD API
// @version         1.0
// @description     A simple CRUD API built with Go, Gin, and GORM
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth
func main() {
	config.LoadEnv()
	database.Connect()

	r := gin.Default()
	routes.RegisterRoutes(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}
