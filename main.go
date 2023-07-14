package main

import (
	"fmt"
	"go-middleware/middleware"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	loggingMiddleware := middleware.Logging()

	http.ListenAndServe(":8080", loggingMiddleware(mux))
}
