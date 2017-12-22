package httplogger

import (
	"net/http"
	"net/url"
	"time"
)

type writer struct {
	http.ResponseWriter
	status int
	size   int
}

func (res *writer) WriteHeader(status int) {
	res.status = status

	res.ResponseWriter.WriteHeader(status)
}

func (res *writer) Write(bytes []byte) (int, error) {
	size, err := res.ResponseWriter.Write(bytes)

	res.size += size

	return size, err
}

type Request struct {
	Request  *http.Request
	Response http.ResponseWriter

	IP     string
	Method string
	URL    *url.URL
	Status int
	Size   int
	Time   time.Duration
}

func Wrap(handler http.HandlerFunc, callback func(*Request)) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		writer := &writer{res, 200, 0}
		// copy URL because it may be modified by http.StripPrefix
		url := req.URL
		start := time.Now()

		handler(writer.ResponseWriter, req)
		callback(&Request{
			Request:  req,
			Response: res,

			IP:     req.RemoteAddr,
			Method: req.Method,
			URL:    url,
			Status: writer.status,
			Size:   writer.size,
			Time:   time.Since(start),
		})
	}
}
