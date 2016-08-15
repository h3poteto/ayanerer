package controllers

import (
	"net/http"

	"github.com/flosch/pongo2"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type Sessions struct {
}

func (u *Sessions) New(c echo.Context) error {
	tpl, err := pongo2.DefaultSet.FromFile("sessions/new.html.tpl")
	if err != nil {
		_ = errors.Wrap(err, "template error")
		return c.HTML(http.StatusInternalServerError, "Internal Server Error")
	}
	tpl.ExecuteWriter(pongo2.Context{"title": "Ayanerer"}, c.Response())
	return nil
}
