package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/TobiEiss/freshboard/plugins"
)

func main() {
	router := echo.New()
	router.Use(middleware.Recover())
	router.Use(middleware.Logger())

	// init freeboard
	router.Static("/", "freeboard/")
	router.File("/config.json", "config.json")
	router.File("/", "freeboard/index.html")

	router.GET("/icscalendar", plugins.UpcomingEvents)

	// start router
	log.Println("start router..")
	router.Logger.Fatal(router.Start(":8080"))
}
