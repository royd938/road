package road

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ContentType = "Content-Type"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (r *Context) TEXT(code int, text string)  {
	r.w.Header().Set(ContentType, "text/plain")
	r.w.WriteHeader(code)
	fmt.Fprint(r.w, text)
}

func (r *Context) HTML(code int, html string)  {
	r.w.Header().Set(ContentType, "text/html")
	r.w.WriteHeader(code)
	fmt.Fprint(r.w, html)
}

func (r *Context) JSON(code int, i interface{})  {
	r.w.Header().Set(ContentType, "application/json")
	r.w.WriteHeader(code)
	bytes, _ := json.Marshal(i)
	fmt.Fprint(r.w, string(bytes))
}
