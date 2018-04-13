package controller

import (
	"net/http"

	"github.com/go-examples/swaggo-sample/account"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

// AccountController account controller
type AccountController struct{}

// NewAccountController mount account controller
func NewAccountController(e *echo.Echo) {
	handler := &AccountController{}

	e.GET("/accounts/:id", handler.Show)
	e.POST("/accounts", handler.Create)
}

// Show show account
// @Summary Show a account
// @Description get account by ID
// @Accept json
// @Produce json
// @Param id path uuid.UUID true "Account ID"
// @Success 200 {object} account.Account
// @Router /accounts/{id} [get]
func (c *AccountController) Show(ctx echo.Context) error {
	res := &account.Account{
		ID:       uuid.NewV4(),
		Email:    "sample@example.com",
		Password: "asdf0987",
	}

	return ctx.JSON(http.StatusOK, res)
}

// Create create account
// @Summary Create a account
// @Description create account
// @Accept json
// @Produce json
// @Param account body account.Payload true "account"
// @Success 201 {object} account.Payload
// @Failere 422 {object} error
// @Router /accounts [post]
func (c *AccountController) Create(ctx echo.Context) error {
	payload := account.Payload{}
	err := ctx.Bind(&payload)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	return ctx.JSON(http.StatusCreated, &payload)
}
