package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"protected-service/middleware"
	"protected-service/services"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(
	service *services.OrderService,
) *OrderHandler {

	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) GetOrders(
	w http.ResponseWriter,
	r *http.Request,
) {

	userIDString :=
		r.Context().Value(
			middleware.UserIDKey,
		).(string)

	userID, err :=
		strconv.ParseInt(
			userIDString,
			10,
			64,
		)

	if err != nil {

		http.Error(
			w,
			"invalid user id",
			http.StatusBadRequest,
		)

		return
	}

	orders, err :=
		h.service.GetOrders(
			r.Context(),
			userID,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(
		orders,
	)
}
