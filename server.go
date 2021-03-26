package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/safecornerscoffee/echo-twitter/handler"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Initialize Handler
	h := &handler.Handler{}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/posts", h.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
