package road

import (
	"log"
	"net/http"
)

type Handler func(c *Context)

type RouteEntry struct {
	Method string
    Path string
    Regex string
	Handler Handler
}

type Router struct {
	RouteEntries []*RouteEntry
}

func New() *Router {
	router := &Router{}
	http.HandleFunc("/", router.resolveHandler)
	return router
}

func (r *Router) GET(path string, handler Handler) {
	r.RouteEntries = append(r.RouteEntries, &RouteEntry{Method: "GET", Path: path, Handler: handler})
}

func (r *Router) POST(path string, handler Handler) {
	r.RouteEntries = append(r.RouteEntries, &RouteEntry{Method: "POST", Path: path, Handler: handler})
}

func (r *Router) Start(path string) error {
	log.Printf("Listen on %s\n", path)
	return http.ListenAndServe(":5000", nil)
}

func (r *Router) resolveHandler(w http.ResponseWriter, req *http.Request) {
	context := &Context{w: w, r: req}
	var routeEntry *RouteEntry
	for _, re := range r.RouteEntries {
		if re.Path == req.RequestURI {
			routeEntry = re
			break
		}
	}
	if routeEntry != nil {
		if routeEntry.Method == req.Method {
			routeEntry.Handler(context)
		} else {
			context.JSON(405, "Method Not Allowed")
		}
	} else {
		context.JSON(404, "Not Found")
	}
}
