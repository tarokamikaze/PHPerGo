package main

import (
	"app/lib"
	"app/service"
	"fmt"
)

var envPrefix string = "TEST_"

func main() {
	container := lib.Bootstrap(envPrefix)
	service := container.MustGet("service").(service.Service)

	fmt.Println(service.Call(container.Config["PARAM"].(string) + "!"))
}