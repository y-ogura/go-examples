package main

import (
	"io/ioutil"

	"github.com/go-examples/goa/response_file/app"
	"github.com/goadesign/goa"
	"github.com/k0kubun/pp"
)

// FileController implements the file resource.
type FileController struct {
	*goa.Controller
}

// NewFileController creates a file controller.
func NewFileController(service *goa.Service) *FileController {
	return &FileController{Controller: service.NewController("FileController")}
}

// File runs the file action.
func (c *FileController) File(ctx *app.FileFileContext) error {
	// FileController_File: start_implement

	// Put your logic here
	f, err := ioutil.ReadFile("./public/sample.xlsx")
	if err != nil {
		return err
	}
	pp.Println(f)
	ctx.ResponseData.Header().Set("Content-Disposition", "attachment; filename=sample.xlsx")
	res := f
	// ctx.ResponseData.Write(res)
	// FileController_File: end_implement
	return ctx.OK(res)
}
