package main

import (
	"IoTSer/controllers"
	"IoTSer/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

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

	router.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/estimations", estimationHandler.CreateEstimation).Methods("POST")
	router.HandleFunc("/cars", carHandler.CreateCar).Methods("POST")
	router.HandleFunc("/sensors", sensorHandler.CreateSensor).Methods("POST")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
