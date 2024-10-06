package goware

import (
	"log"
	"net/http"
)

type GoWare struct {
	logger *log.Logger
}

func (gw *GoWare) SetupLogger(logger *log.Logger) {
	gw.logger = logger
	logger.Println("Logger is set")
}

func (gw *GoWare) Use(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		if gw.logger != nil {
			gw.logger.Println("Request Reveived:", r.Method, r.URL.Path)
		}
		next.ServeHTTP(w, r)
	})
}
