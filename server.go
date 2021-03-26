package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	"github.com/safecornerscoffee/echo-twitter/handler"

	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Database connection
	postgresHost := os.Getenv("POSTGRES_HOST")
	postgresPort := os.Getenv("POSTGRES_PORT")
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDb := os.Getenv("POSTGRES_DB")
	pqInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		postgresHost, postgresPort, postgresUser, postgresPassword, postgresDb)

	db, err := sql.Open("postgres", pqInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for i := 0; i < 360; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Initialize Handler
	h := &handler.Handler{DB: db}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/posts", h.CreatePost)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
