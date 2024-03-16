package main

import (
	"github.com/lpernett/godotenv"

	"github.com/2gbeh/drygin-api/configs/db_config"
	"github.com/2gbeh/drygin-api/configs/gin_config"

	// "github.com/2gbeh/drygin-api/configs/gin_config"
	"github.com/2gbeh/drygin-api/controllers"

	// _ "github.com/2gbeh/drygin-api/docs"
	"github.com/2gbeh/drygin-api/models"
	"github.com/2gbeh/drygin-api/repositories"
	"github.com/2gbeh/drygin-api/routes"
	"github.com/2gbeh/drygin-api/utils"
)

// @title 	Drygin
// @version	1.0
// @description Microsoft Northwind Database REST API

// @host 	localhost:8000
// @BasePath /api

func init(){
  // Load env file
  err := godotenv.Load()
  utils.LogFatal(err, "Load .env file failed.")
}

func main() {
	// Connect database
	db, tb := db_config.Connect()

	// Run migrations
	db.Table(tb.Tags).AutoMigrate(&models.Tag{})

	// Get repositories
	tagRepository := repositories.UseTagRpository(db, tb)

	// Get controllers
	tagController := controllers.UseTagController(tagRepository)

	// Get routes
	routes := routes.Routes(tagController)

	// Start server
	gin_config.Start(routes)
}
