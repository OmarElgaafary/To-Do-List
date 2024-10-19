package server

import (
	"log"
	"net/http"
	"todo/internal/handlers"

	"github.com/labstack/echo/v4"
)

func RunServer(h *handlers.ToDoHandler) {
	server := echo.New()
	CreateToDoRoutes(server, h)
	err := server.Start(":8080")
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func CreateToDoRoutes(e *echo.Echo, h *handlers.ToDoHandler) {
	todoGroup := e.Group("/todo")
	todoGroup.POST("/Create", h.CreateToDoHandler)
	todoGroup.GET("/:id", h.GetToDoByIDHandler)
	// todoGroup.GET("/:id", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "hello todo from omar")
	// })
}
