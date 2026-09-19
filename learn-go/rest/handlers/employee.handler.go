package handlers

import (
	"encoding/json"
	"net/http"

	"example.com/first-app/rest/models"
	"example.com/first-app/rest/service"
)

type EmployeeHandler struct {
	Service service.EmployeeService
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees, err := h.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func (h *EmployeeHandler) CreateNewEmployee(w http.ResponseWriter, r *http.Request) {
	var Employee models.Employee
	err := json.NewDecoder(r.Body).Decode(&Employee)
	if err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}
	createdEmployee, err := h.Service.CreateEmployee(Employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEmployee)
}
