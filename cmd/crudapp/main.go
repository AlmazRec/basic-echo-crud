package main

import (
	"awesomeProject3/config"
	"awesomeProject3/internal/database"
	"awesomeProject3/internal/handlers"
	"awesomeProject3/internal/repositories"
	"awesomeProject3/internal/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	cfg := &config.Config{}

	cfg, err := cfg.Load("../../config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDB(cfg)

	postRepo := repositories.NewPostRepository(db)
	postHandler := handlers.NewPostHandler(postRepo)
	h := handlers.NewHandler(postHandler)
	routes.InitRoutes(e, *h)

	err = e.Start(":" + cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
