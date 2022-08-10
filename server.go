package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	ProjectID  int
}

var employees = []Employee{
	{1, "Gaurav", "LnD", 1001},
	{2, "Jose", "Cloud", 1002},
	{3, "Prabhakar", "SRE", 10003},
}

func EmployeesIndexHandler(w http.ResponseWriter, req *http.Request) {
	// fmt.Fprintln(w, employees)
	json.NewEncoder(w).Encode(employees)
}

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		msg := "Hello, World"

		// fmt.Printf("%T\n", w)

		// w.Write([]byte(msg))
		fmt.Fprintln(w, msg)
	})

	r.HandleFunc("/employees", EmployeesIndexHandler)

	http.ListenAndServe("localhost:8000", r)
}
