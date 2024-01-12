package models

import "database/sql"

type Car struct {
	ID        int    `json:"id"`
	Number    string `json:"number"`
	Equipment string `json:"equipment"`
	Credit    string `json:"credit"`
}

type CarModel struct {
	DB *sql.DB
}

func NewCarModel(DB *sql.DB) *CarModel {
	return &CarModel{DB: DB}
}

func (m *CarModel) All() ([]Car, error) {
	rows, err := m.DB.Query("SELECT * FROM Car")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var carinfos []Car
	for rows.Next() {
		var carinfo Car
		if err := rows.Scan(
			&carinfo.ID,
			&carinfo.Number,
			&carinfo.Equipment,
			&carinfo.Credit); err != nil {
			return nil, err
		}
		carinfos = append(carinfos, carinfo)
	}

	return carinfos, nil
}

func (m *CarModel) Insert(id int, number string, equipment string, credit string) (int, error) {
	_, err := m.DB.Exec("INSERT INTO Car (id, number, equipment, credit) VALUES (?, ?, ?, ?)", id, number, equipment, credit)
	if err != nil {
		return 0, err
	}

	return id, nil
}
