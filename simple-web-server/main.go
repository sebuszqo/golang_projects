package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Now())
		next.ServeHTTP(w, r)
	})
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	router := http.NewServeMux()
	router.Handle("/", loggingMiddleware(fileServer))
	router.Handle("/form", loggingMiddleware(http.HandlerFunc(formHandler)))
	router.Handle("/hello", loggingMiddleware(http.HandlerFunc(helloHandler)))

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
