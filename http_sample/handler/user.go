package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-examples/http_sample/repository"
	"github.com/go-examples/http_sample/usecase"
)

// UserHandler user http handler
func UserHandler(w http.ResponseWriter, r *http.Request) {
	conn := &sql.DB{}
	userRepo := repository.NewUserRepository(conn)
	userUsecase := usecase.NewUserUsecase(userRepo)
	c := NewUserController(userUsecase, w, r)
	switch r.Method {
	case "GET":
		c.List()
	case "POST":
		c.Create()
	default:
		fmt.Println("404 error")
	}
}

// userController user controller
type userController struct {
	userUsecase usecase.UserUsecase
	Context     Context
}

// NewUserController mount user controller
func NewUserController(us usecase.UserUsecase, w http.ResponseWriter, r *http.Request) UserController {
	return &userController{
		userUsecase: us,
		Context: Context{
			w: w,
			r: r,
		},
	}
}

// UserController user controller interface
type UserController interface {
	List()
	Create()
}

// List get user list controller
func (u *userController) List() {
	res, err := u.userUsecase.List()
	encoder := json.NewEncoder(u.Context.w)
	if err != nil {
		u.Context.w.WriteHeader(http.StatusBadGateway)
		encoder.Encode(err)
		return
	}
	u.Context.w.WriteHeader(http.StatusOK)
	encoder.Encode(res)
	return
}

// Create create new user controller
func (u *userController) Create() {
	return
}
