package controllers

import (
	"IoTSer/models"
	"encoding/json"
	"net/http"
)

type SensorController struct {
	models *models.SensorModel
}

func NewSensorController(models *models.SensorModel) *SensorController {
	return &SensorController{models: models}
}

func (h *SensorController) GetSensors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	sencors, err := h.models.All()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(sencors)
}

func (h *SensorController) CreateSensor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sencor models.Sensor
	if err := json.NewDecoder(r.Body).Decode(&sencor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.models.Insert(sencor.ID, sencor.Acceleration_x, sencor.Acceleration_y, sencor.Acceleration_z)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sencor.ID = id
	json.NewEncoder(w).Encode(sencor)
}
