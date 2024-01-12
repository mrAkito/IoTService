package models

import "database/sql"

type Sensor struct {
	ID             int     `json:"id"`
	Acceleration_x float64 `json:"acceleration_x"`
	Acceleration_y float64 `json:"acceleration_y"`
	Acceleration_z float64 `json:"acceleration_z"`
}

type SensorModel struct {
	DB *sql.DB
}

func NewSensorModel(DB *sql.DB) *SensorModel {
	return &SensorModel{DB: DB}
}

func (m *SensorModel) All() ([]Sensor, error) {
	rows, err := m.DB.Query("SELECT * FROM Sensor")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sensors []Sensor
	for rows.Next() {
		var sensor Sensor
		if err := rows.Scan(
			&sensor.ID,
			&sensor.Acceleration_x,
			&sensor.Acceleration_y,
			&sensor.Acceleration_z); err != nil {
			return nil, err
		}
		sensors = append(sensors, sensor)
	}

	return sensors, nil
}

func (m *SensorModel) Insert(id int, acceleration_x float64, acceleration_y float64, acceleration_z float64) (int, error) {
	_, err := m.DB.Exec("INSERT INTO Sensor (id, acceleration_x, acceleration_y, acceleration_z) VALUES (?, ?, ?, ?)", id, acceleration_x, acceleration_y, acceleration_z)
	if err != nil {
		return 0, err
	}

	return id, nil
}
