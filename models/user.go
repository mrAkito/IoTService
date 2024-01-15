package models

import (
	"database/sql"
	"time"
)

// Globalmap
var matching_Map = map[string]string{}

type Message struct {
	Message string `json:"message"`
}

type UserKind struct {
	Name string `json:"username"`
	kind int    `json:"kind"`
}

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	Collab_service string `json:"collab_service"`
	Login          bool   `json:"login"`
	Birthday       string `json:"birthday"`
	Payment_method string `json:"payment_method"`
	Phone_number   string `json:"phone_number"`
	Kind           bool   `json:"kind"`
}

type UserModel struct {
	DB *sql.DB
}

func NewUserModel(DB *sql.DB) *UserModel {
	return &UserModel{DB: DB}
}

func (m *UserModel) All() ([]User, error) {
	rows, err := m.DB.Query("SELECT * FROM User")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userinfos []User
	for rows.Next() {
		var userinfo User
		if err := rows.Scan(
			&userinfo.ID,
			&userinfo.Username,
			&userinfo.Password,
			&userinfo.Collab_service,
			&userinfo.Login,
			&userinfo.Birthday,
			&userinfo.Payment_method,
			&userinfo.Phone_number,
			&userinfo.Kind); err != nil {
			return nil, err
		}
		userinfos = append(userinfos, userinfo)
	}

	return userinfos, nil
}

func (m *UserModel) GetOne(id int) ([]User, error) {
	rows, err := m.DB.Query("SELECT * FROM User WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userinfos []User
	for rows.Next() {
		var userinfo User
		if err := rows.Scan(
			&userinfo.ID,
			&userinfo.Username,
			&userinfo.Password,
			&userinfo.Collab_service,
			&userinfo.Login,
			&userinfo.Birthday,
			&userinfo.Payment_method,
			&userinfo.Phone_number,
			&userinfo.Kind); err != nil {
			return nil, err
		}
		userinfos = append(userinfos, userinfo)
	}

	return userinfos, nil
}

func (m *UserModel) Insert(username string, password string, collab_service string, login bool, birthday string, payment_method string, phone_number string, kind bool) (int, error) {
	birthdayFormat, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		return 0, err
	}

	result, err := m.DB.Exec("INSERT INTO User (username, password, collab_service, login, birthday, payment_method, phone_number, kind) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", username, password, collab_service, login, birthdayFormat, payment_method, phone_number, kind)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *UserModel) Login(username string, password string) (Message, error) {
	row, err := m.DB.Query("SELECT kind FROM User WHERE username = ? AND password = ?", username, password)
	if err != nil {
		return Message{"failed"}, err
	}

	_, err = m.DB.Exec("UPDATE User SET login = 1 WHERE username = ? AND password = ?", username, password)
	if err != nil {
		return Message{"failed"}, err
	}
	defer row.Close()

	var kind string
	if row.Next() {
		if err := row.Scan(&kind); err != nil {
			return Message{"failed"}, err
		}
	}

	if kind == "false" {
		kind = "driver"
	} else {
		kind = "costomer"
	}

	return Message{kind}, nil
}

func (m *UserModel) Logout(id int) (Message, error) {
	row, err := m.DB.Query("SELECT kind FROM User WHERE id = ?", id)
	if err != nil {
		return Message{"failed"}, err
	}

	_, err = m.DB.Exec("UPDATE User SET login = 0 WHERE id = ?", id)
	if err != nil {
		return Message{"failed"}, err
	}
	defer row.Close()

	var kind string
	if row.Next() {
		if err := row.Scan(&kind); err != nil {
			return Message{"failed"}, err
		}
	}

	if kind == "false" {
		kind = "costomer"
	} else {
		kind = "driver"
	}

	return Message{kind}, nil
}

// func (m *UserModel) customer() (Message, error) {
// 	row, err := m.DB.Query("SELECT username, kind FROM User WHERE login = 1 AND kind = 0")
// 	if err != nil {
// 		return Message{"failed"}, err
// 	}

// 	defer row.Close()

// 	if len(row) == 0 {
// 		// 再帰する
// 		// customer(w, r)
// 	} else {
// 		// 乱数を生成
// 		rand.Seed(time.Now().UnixNano())
// 		index := rand.Intn(len(result))
// 		driver := result[index].Name
// 		matching_Map[driver] = name
// 		m := Message{driver}
// 	}
// }
