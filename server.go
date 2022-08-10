package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Employee struct {
	ID         int    `json:"-"` // Struct tags
	Name       string `json:"name"`
	Department string `json:"speciality"`
	ProjectID  int    `json:"project"`
}

// func (e Employee) MarshalJSON() ([]byte, error) {
// 	jsonString := fmt.Sprintf(`{"name": "%s", "speciality": "%s", "project": %d}`, e.Name, e.Department, e.ProjectID)

// 	return []byte(jsonString), nil
// }

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Jose", "Cloud", 1002},
	{3, "Prabhakar", "SRE", 10003},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

func EmployeeCreateHandler(w http.ResponseWriter, req *http.Request) {
	var newEmployee Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	newEmployee.ID = len(employees) + 1

	employees = append(employees, newEmployee)

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}

// func EmployeesHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "POST" {
// 		EmployeeCreateHandler(w, req)
// 	} else if req.Method == "GET" {
// 		EmployeesIndexHandler(w, req)
// 	} else {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		return
// 	}
// }

func EmployeeShowHandler(w http.ResponseWriter, req *http.Request) {
	employeeID := mux.Vars(req)["id"]

	empID, _ := strconv.Atoi(employeeID)

	emp := employees[empID-1]

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func loggingMiddleware(next http.Handler) http.Handler {
	f := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		// if authorized {
		next.ServeHTTP(w, req)
		// } else {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	return
		// }

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

	http.ListenAndServe("localhost:8000", loggingMiddleware(r))
}
