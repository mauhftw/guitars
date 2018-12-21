package main

import (
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/mauhftw/go-guitarists/controllers"
	"github.com/mauhftw/go-guitarists/helpers"
)

func main() {
	e := echo.New()

	// Initialize uptime
	var StartTime = time.Now()
	helpers.StartTime = StartTime

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/version", controllers.GetVersion)
	e.GET("/health", controllers.GetHealth)
	e.GET("/users", controllers.GetUsers)
	e.GET("/hellos", controllers.GetHellos)

	// Start server
	e.Logger.Fatal(e.Start(":1337"))
}
