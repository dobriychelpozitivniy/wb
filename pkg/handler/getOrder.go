package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) getOrder(c echo.Context) error {
	id := c.Param("id")
	order, err := h.s.GetMessage(id)
	if err != nil {
		fmt.Println(err)
		return c.String(404, err.Error())
	}

	return c.Render(http.StatusOK, "template.html", order)
}
