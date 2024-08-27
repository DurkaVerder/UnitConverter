package handler

import (
	"UnitConverter/interval/conv"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Run(e *echo.Echo) {

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t
	e.Static("/static", "static")
	e.Use(middleware.Logger())

	e.GET("/unitconverter/length/:val/:from/:to/:result", func(c echo.Context) error {

		return c.String(http.StatusOK, "OK")
	})

	e.POST("/unitconverter/length", func(c echo.Context) error {
		value, err := strconv.ParseFloat(c.FormValue("value"), 64)
		if err != nil {
			return err
		}

		convertFrom := c.FormValue("ConvertFrom")
		convertTo := c.FormValue("ConvertTo")
		result := conv.ConverterLength(convertFrom, convertTo, value)
		return c.Redirect(http.StatusOK, fmt.Sprintf("/unitconverter/length/:%v/:%v:/%v/:%v", c.FormValue("value"), convertFrom, convertTo, result))
	})

	e.Start(":2222")
}
