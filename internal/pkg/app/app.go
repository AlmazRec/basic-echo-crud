package app

import (
	"awesomeProject3/config"
	"awesomeProject3/internal/database"
	"awesomeProject3/internal/delivery/http"
	"awesomeProject3/internal/delivery/http/routes"
	"awesomeProject3/internal/repositories"
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

type App struct {
	Echo *echo.Echo
	Cfg  *config.Config
	DB   *sql.DB
}

func New(path string) *App {
	e := echo.New()
	e.Use(middleware.Logger())

	log.Println("Start load cfg")
	cfg := &config.Config{}
	cfg, err := cfg.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDB(cfg)

	postRepo := repositories.NewPostRepository(db)
	postHandler := http.NewPostHandler(postRepo)
	h := http.NewHandler(postHandler)
	routes.InitRoutes(e, *h)

	return &App{
		Echo: e,
		Cfg:  cfg,
		DB:   db,
	}
}

func (a *App) Start() {
	err := a.Echo.Start(":" + a.Cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
	}
}
