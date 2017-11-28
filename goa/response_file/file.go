package main

import (
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/go-examples/goa/response_file/app"
	"github.com/goadesign/goa"
	"github.com/tealeg/xlsx"
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
	// make excel file.
	var fileName = "sample.xlsx"
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	exec.Command("rm -rf ./public/" + fileName)
	file = xlsx.NewFile()
	sheet, err := file.AddSheet("sample_sheet")
	if err != nil {
		return err
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "create excel file!"
	err = file.Save("./public/" + fileName)
	if err != nil {
		return err
	}
	log.Println("make file")

	// get excel file
	res, err := ioutil.ReadFile("./public/" + fileName)
	if err != nil {
		return err
	}
	// set response header for content disposition
	ctx.ResponseData.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	log.Println("download excel file")
	defer exec.Command("rm -rf ./public/" + fileName)
	return ctx.OK(res)
}
