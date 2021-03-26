package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/safecornerscoffee/echo-twitter/model"
)

func (h *Handler) CreatePost(c echo.Context) (err error) {
	// u := &model.User{}
	p := &model.Post{}
	if err = c.Bind(p); err != nil {
		return
	}
	return c.JSON(http.StatusCreated, p)
}
