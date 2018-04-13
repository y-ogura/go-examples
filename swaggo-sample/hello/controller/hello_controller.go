package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloController hello controller
type HelloController struct{}

// NewHelloController mount hello controller
func NewHelloController(e *echo.Echo) {
	handler := &HelloController{}

	e.GET("/hello", handler.Hello)
}

// Hello hello world
// @Summary hello world
// @Description show hello
// @Accept  json
// @Produce  json
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /hello [get]
func (c *HelloController) Hello(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "hello, world!")
}
