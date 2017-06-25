package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/reujab/httplogger"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello, world!"))
	}).Methods("GET")
	http.ListenAndServe(":8080", httplogger.Wrap(router.ServeHTTP, func(req *httplogger.Request) {
		if !strings.HasPrefix(req.URL.Path, "/socket.io/") {
			log.Println(req.IP, req.Method, req.URL, req.Status, req.Time)
		}
	}))
}
