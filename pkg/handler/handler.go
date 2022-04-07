package handler

import (
	"text/template"
	"wbl0/pkg/service"

	"github.com/labstack/echo/v4"
)

// import "github.com/labstack/echo/v4"

type Handler struct {
	s *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{s: service}
}

func (h *Handler) initRoutes() *echo.Echo {
	e := echo.New()

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = renderer

	e.GET("/orders/:id", h.getOrder)

	return e
}

func (h *Handler) StartServer() {
	e := h.initRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}
