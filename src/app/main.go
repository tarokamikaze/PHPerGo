package main

import (
	"github.com/fgrosse/goldi"
	"app/lib"
	"github.com/fgrosse/goldi/validation"
	"app/service"
	"fmt"
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
	service := container.MustGet("service").(service.Service)

	fmt.Println(service.Call("test!"))
}