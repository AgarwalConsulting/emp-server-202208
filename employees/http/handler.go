package http

import (
	"algogrit.com/emp-server/employees/service"
	"github.com/gorilla/mux"
)

type EmployeeHandler struct {
	*mux.Router // Embedding => Inheritance {ServeHTTP}
	// Router *mux.Router

	svcV1 service.EmployeeService
}

func (e *EmployeeHandler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/v1/employees", e.IndexV1).Methods("GET")
	r.HandleFunc("/v1/employees", e.CreateV1).Methods("POST")

	r.HandleFunc("/v1/employees/{id}", e.ShowV1).Methods("GET")

	e.Router = r
}

func New(svcV1 service.EmployeeService) EmployeeHandler {
	h := EmployeeHandler{svcV1: svcV1}

	h.SetupRoutes(mux.NewRouter())

	return h
}
