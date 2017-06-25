package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/reujab/httplogger"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	})
	http.HandleFunc("/404", func(res http.ResponseWriter, req *http.Request) {
		http.NotFound(res, req)
	})
	http.ListenAndServe(":8080", httplogger.Wrap(http.DefaultServeMux.ServeHTTP, func(req *httplogger.Request) {
		if !strings.HasPrefix(req.URL.Path, "/socket.io/") {
			log.Println(req.IP, req.Method, req.URL, req.Status, req.Time)
		}
	}))
}
