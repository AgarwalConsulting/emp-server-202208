package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	empHTTP "algogrit.com/emp-server/employees/http"
	"algogrit.com/emp-server/employees/repository"
	"algogrit.com/emp-server/employees/service"
)

func loggingMiddleware(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		next.ServeHTTP(w, req)

		dur := time.Since(begin)

		log.Printf("URL: %s %s | dur: %s", req.Method, req.URL, dur)
	}

	return http.HandlerFunc(f)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World"

		fmt.Fprintln(w, msg)
	})

	var empRepo = repository.NewInMem()
	var empV1 = service.NewV1(empRepo)
	var empHandler = empHTTP.New(empV1)

	empHandler.SetupRoutes(r)

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe("localhost:8000", loggingMiddleware(r))
}
