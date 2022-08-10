package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"algogrit.com/emp-server/entities"
)

func (h *EmployeeHandler) IndexV1(w http.ResponseWriter, req *http.Request) {
	employees, err := h.svcV1.Index()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) CreateV1(w http.ResponseWriter, req *http.Request) {
	var newEmployee entities.Employee

	decoder := json.NewDecoder(req.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	createdEmp, err := h.svcV1.Create(newEmployee)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdEmp)
}

func (h *EmployeeHandler) ShowV1(w http.ResponseWriter, req *http.Request) {
	employeeID := mux.Vars(req)["id"]

	empID, _ := strconv.Atoi(employeeID)

	emp, err := h.svcV1.Show(empID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
