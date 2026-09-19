package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/vignesh/online-ordering/internal/models"
	"github.com/vignesh/online-ordering/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	err = h.service.Create(r.Context(), &order)
	if err != nil {
		http.Error(w, "failed to create order", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid order ID", http.StatusBadRequest)
		return
	}
	order, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(order)
}
