package main

import (
	"./controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func Routes(e *echo.Echo) {
	root := controllers.Root{}
	e.Get("/", root.Index)
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("public"))

	Routes(e)

	e.Run(standard.New(":9090"))
}
