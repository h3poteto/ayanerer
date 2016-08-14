package controllers

import (
	"github.com/labstack/echo"
)

type Root struct {
}

func (u *Root) Index(c echo.Context) error {
	return nil
}
