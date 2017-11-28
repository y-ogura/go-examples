package main

import (
	"io/ioutil"
	"log"

	"github.com/go-examples/goa/response_file/app"
	"github.com/goadesign/goa"
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
	// get excel file
	res, err := ioutil.ReadFile("./public/sample.xlsx")
	if err != nil {
		return err
	}
	// set response header for content disposition
	ctx.ResponseData.Header().Set("Content-Disposition", "attachment; filename=sample.xlsx")
	log.Println("download excel file")
	return ctx.OK(res)
}
