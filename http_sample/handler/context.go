package handler

import "net/http"

// Context context
type Context struct {
	w http.ResponseWriter
	r *http.Request
}
