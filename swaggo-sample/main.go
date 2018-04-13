package main

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
// @BasePath /

import (
	accountC "github.com/go-examples/swaggo-sample/account/controller"
	helloC "github.com/go-examples/swaggo-sample/hello/controller"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"

	_ "github.com/go-examples/swaggo-sample/docs" // docs in generate by Swag CLI
)

func main() {
	e := echo.New()

	// mount controllers
	helloC.NewHelloController(e)
	accountC.NewAccountController(e)

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
