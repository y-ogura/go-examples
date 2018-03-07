// @APIVersion 1.0.0
// @APITitle echo-swagger
// @APIDescription echo-swagger sample api
// @BasePath http://localhost:8080/swagger-ui
// @SubApi hello [/hello]

package main

import (
	helloC "github.com/go-examples/echo-swagger/hello/controller"
	swaggerui "github.com/go-examples/echo-swagger/swagger-ui"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	helloC.NewHelloController(e)

	swaggerui.NewSwaggerController(e)

	e.Logger.Fatal(e.Start(":8080"))
}
