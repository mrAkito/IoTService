package controllers

import (
	"IoTSer/models"
	"encoding/json"
	"net/http"
)

type EstimationController struct {
	models *models.EstimationModel
}

func NewEstimationController(models *models.EstimationModel) *EstimationController {
	return &EstimationController{models: models}
}

func (h *EstimationController) GetEstimations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	estimations, err := h.models.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(estimations)
}

func (h *EstimationController) CreateEstimation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var estimation models.Estimation
	err := json.NewDecoder(r.Body).Decode(&estimation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id, err := h.models.Insert(estimation.ID, estimation.Estimation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	estimation.ID = id
	json.NewEncoder(w).Encode(estimation)
}
