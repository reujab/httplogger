# HTTPLogger [![Documentation](https://godoc.org/github.com/reujab/httplogger?status.svg)](https://godoc.org/github.com/reujab/httplogger)
A Golang library that logs HTTP requests using custom logic.

## Examples

### Standard library
```go
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
```

### Gorilla Mux
```go
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
```

### Log
```
2017/06/25 14:47:49 [::1]:45616 GET / 200 5.721µs
2017/06/25 14:47:51 [::1]:45616 GET /404 404 48.588µs
```
