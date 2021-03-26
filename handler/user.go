package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/safecornerscoffee/echo-twitter/model"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	return c.JSON(http.StatusOK, u)
}
