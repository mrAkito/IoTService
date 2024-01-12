package controllers

import (
	"IoTSer/models"
	"encoding/json"
	"net/http"
)

type CarController struct {
	models *models.CarModel
}

func NewCarController(models *models.CarModel) *CarController {
	return &CarController{models: models}
}

func (h *CarController) GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cars, err := h.models.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(cars)
}

func (h *CarController) CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := h.models.Insert(car.ID, car.Number, car.Equipment, car.Credit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	car.ID = id
	json.NewEncoder(w).Encode(car)
}
