package goware

import (
	"fmt"
	"net/http"
)

func MyMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Reveived:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Println("Respnse Sent")
	})
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", MyMiddleware(http.HandlerFunc(mainHandler)))

	http.ListenAndServe(":8080", mux)

}
