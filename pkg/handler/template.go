package handler

import (
	"io"
	"text/template"
	"wbl0/pkg/model"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *TemplateRenderer) ToEchoMap(msg *model.Message) *echo.Map {
	return &echo.Map{
		"order_uid": "test",
	}
}
