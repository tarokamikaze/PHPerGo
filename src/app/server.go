package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/fgrosse/goldi"
	"app/lib"
	"github.com/fgrosse/goldi/validation"
	"app/service"
)

func main() {
	registry := goldi.NewTypeRegistry()
	lib.RegisterTypes(registry)

	config := map[string]interface{}{
		"some_parameter": "Hello World",
		"timeout":        42.7,
	}
	container := goldi.NewContainer(registry, config)
	validator := validation.NewContainerValidator()
	validator.MustValidate(container)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		service := container.MustGet("service").(service.Service)
		return c.String(http.StatusOK, service.Call("test!"))
	})
	e.Logger.Fatal(e.Start(":1323"))
}