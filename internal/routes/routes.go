package routes

import (
	"awesomeProject3/internal/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, h handlers.Handler) {
	PostRoutes(e, h)
}

func PostRoutes(e *echo.Echo, h handlers.Handler) {
	e.GET("/posts", h.PostHandler.GetPosts)
	e.GET("/posts/:id", h.PostHandler.GetPost)
	e.POST("/posts", h.PostHandler.StorePost)
	e.PUT("/posts/:id", h.PostHandler.UpdatePost)
	e.DELETE("/posts/:id", h.PostHandler.DeletePost)
}
