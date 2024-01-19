package main

import (
	"IoTSer/controllers"
	"IoTSer/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set headers
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			// Otherwise, pass the request to the next middleware in the chain
			next.ServeHTTP(w, r)
		})
	}

	router.Use(corsMiddleware)

	db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/IoTSer_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//ユーザ情報取得
	userModel := models.NewUserModel(db)
	userHandler := controllers.NewUserController(userModel)

	//車情報取得
	carModel := models.NewCarModel(db)
	carHandler := controllers.NewCarController(carModel)

	//評価情報取得
	estimationModel := models.NewEstimationModel(db)
	estimationHandler := controllers.NewEstimationController(estimationModel)

	//センサ情報取得
	sensorModel := models.NewSensorModel(db)
	sensorHandler := controllers.NewSensorController(sensorModel)

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/cars", carHandler.GetCars).Methods("GET")
	router.HandleFunc("/estimations", estimationHandler.GetEstimations).Methods("GET")
	router.HandleFunc("/sensors", sensorHandler.GetSensors).Methods("GET")

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST", "OPTIONS")
	router.HandleFunc("/estimations", estimationHandler.CreateEstimation).Methods("POST", "OPTIONS")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST", "OPTIONS")
	router.HandleFunc("/sensors", sensorHandler.CreateSensor).Methods("POST", "OPTIONS")

	http.HandleFunc("/customer", customer)
	http.HandleFunc("/driver", driver)

	router.HandleFunc("/login", userHandler.Login).Methods("PUT", "OPTIONS")
	router.HandleFunc("/logout", userHandler.Logout).Methods("PUT", "OPTIONS")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

var matching_Map = map[string]string{}

type Message struct {
	Message string `json:"message"`
}

type User struct {
	Name string `json:"username"`
	kind int    `json:"kind"`
}

func getquery(Query string) []User {
	// データベースに接続
	db, err := sql.Open("mysql", "user:userpassword@tcp(localhost:3306)/IoTSer_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// データベースに接続できるか確認
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// データを取得
	rows, err := db.Query(Query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 取得したデータを戻り値にする
	var result []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Name, &user.kind)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, user)
	}
	return result
}

func customer(w http.ResponseWriter, r *http.Request) {
	// paramsから値を取得
	params := r.URL.Query()
	name := params.Get("name")
	Query := "SELECT username,kind FROM user WHERE login = 1 AND kind = 0"
	result := getquery(Query)
	if len(result) == 0 {
		// 再帰する
		// customer(w, r)
	} else {
		// 乱数を生成
		rand.Seed(time.Now().UnixNano())
		index := rand.Intn(len(result))
		driver := result[index].Name
		matching_Map[driver] = name
		m := Message{driver}
		res, err := json.Marshal(m)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}
}

func driver(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	name := params.Get("name")
	// 無限ループ
	for {
		// マッチングマップがNullでないか確認
		if matching_Map[name] != "" {
			// マッチングマップから値を取得
			customer := matching_Map[name]
			m := Message{customer}
			res, err := json.Marshal(m)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(res)
			break
		}
		// 2秒待つ
		time.Sleep(2 * time.Second)
	}
}
