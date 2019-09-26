package http

/**
Create HTTP server using net/http package
*/

import (
	"fmt"
	"net/http"
)

// handler is a an object implementing the http.Handler interface
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// functions serving handlers take a http.ResponseWriter and http.Request as args

func headers(w http.ResponseWriter, req *http.Request) {
	// read all HTTP request headers and echo them into the response body
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// register handlers on server routes
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// call listen and serve witht the port and a handler
	// nil = default router we have set up
	http.ListenAndServe(":8090", nil)
}
