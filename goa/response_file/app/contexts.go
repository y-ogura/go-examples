// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "file": Application Contexts
//
// Command:
// $ goagen
// --design=github.com/go-examples/goa/response_file/design
// --out=$(GOPATH)/src/github.com/go-examples/goa/response_file
// --version=v1.2.0-dirty

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// FileFileContext provides the file file action context.
type FileFileContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewFileFileContext parses the incoming request URL and body, performs validations and creates the
// context used by the file controller file action.
func NewFileFileContext(ctx context.Context, r *http.Request, service *goa.Service) (*FileFileContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := FileFileContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *FileFileContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.ms-excel")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *FileFileContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
