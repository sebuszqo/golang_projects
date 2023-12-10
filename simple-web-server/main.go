package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	getMethod  = "GET"
	postMethod = "POST"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != postMethod {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() err: %v", err), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("POST request successful\n"))

	name := r.PostForm.Get("name")
	address := r.PostForm.Get("address")

	if name == "" || address == "" {
		http.Error(w, "Name and Address are required fields", http.StatusBadRequest)
		return
	}

	const maxNameLength = 50
	const maxAddressLength = 100

	if len(name) > maxNameLength || len(address) > maxAddressLength {
		http.Error(w, "Name or Address exceeds the maximum allowed length", http.StatusBadRequest)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello: %s, from %s", capitalizeFirstLetter(name), capitalizeFirstLetter(address))))
}

func capitalizeFirstLetter(input string) string {
	if input == "" {
		return input
	}

	return strings.ToUpper(string(input[0])) + input[1:]
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != getMethod {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	w.Write([]byte("hello from there\n"))
}

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
