package main

import (
	"os"
	"path/filepath"

	"./controllers"
	"./filters"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func routes(e *echo.Echo) {
	root := controllers.Root{}
	e.Get("/", root.Index)

	sessions := controllers.Sessions{}
	e.Get("/sessions/new", sessions.New)
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Static("public"))

	setupPongo()

	routes(e)

	e.Run(standard.New(":9090"))
}

func setupPongo() {
	root := os.Getenv("ECHOROOT")
	pongo2.DefaultSet = pongo2.NewSet("default", pongo2.MustNewLocalFileSystemLoader(filepath.Join(root, "views")))
	pongo2.RegisterFilter("suffixAssetsUpdate", filters.SuffixAssetsUpdate)
}
