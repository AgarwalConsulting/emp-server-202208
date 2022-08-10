package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/entities"

	"algogrit.com/emp-server/employees/repository"
)

var empRepo = repository.NewInMem()

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	employees, err := empRepo.ListAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := empRepo.Save(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

func EmployeeShowHandler(w http.ResponseWriter, req *http.Request) {
	employeeID := mux.Vars(req)["id"]

	empID, _ := strconv.Atoi(employeeID)

	emp, err := empRepo.FindBy(empID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

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
	// r := http.NewServeMux()
	r := mux.NewRouter()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World"

		fmt.Fprintln(w, msg)
	})

	// r.HandleFunc("/employees", EmployeesHandler)
	r.HandleFunc("/employees", EmployeesIndexHandler).Methods("GET")
	r.HandleFunc("/employees", EmployeeCreateHandler).Methods("POST")

	r.HandleFunc("/employees/{id}", EmployeeShowHandler).Methods("GET")

	log.Println("Starting server on port: 8000...")
	http.ListenAndServe("localhost:8000", loggingMiddleware(r))
}
