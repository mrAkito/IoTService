package models

import "database/sql"

type Estimation struct {
	ID         int    `json:"id"`
	Estimation string `json:"driver_esti"`
}

type EstimationModel struct {
	DB *sql.DB
}

func NewEstimationModel(DB *sql.DB) *EstimationModel {
	return &EstimationModel{DB: DB}
}

func (m *EstimationModel) All() ([]Estimation, error) {
	rows, err := m.DB.Query("SELECT * FROM Estimation")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var estimations []Estimation
	for rows.Next() {
		var estimation Estimation
		if err := rows.Scan(
			&estimation.ID,
			&estimation.Estimation); err != nil {
			return nil, err
		}
		estimations = append(estimations, estimation)
	}

	return estimations, nil
}

func (m *EstimationModel) Insert(id int, estimation string) (int, error) {
	_, err := m.DB.Exec("INSERT INTO Estimation (id, driver_esti) VALUES (?, ?)", id, estimation)
	if err != nil {
		return 0, err
	}

	return id, nil
}
