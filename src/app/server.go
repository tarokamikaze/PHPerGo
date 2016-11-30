package main

import (
	"net/http"

	"github.com/labstack/echo"
	"app/lib"
	"app/service"
)
var envPrefix string = "TEST_"
func main() {
	container := lib.Bootstrap(envPrefix)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		service := container.MustGet("service").(service.Service)
		str := container.Config["PARAM"].(string) + "!"
		return c.String(http.StatusOK, service.Call(str))
	})
	e.Logger.Fatal(e.Start(":1323"))
}