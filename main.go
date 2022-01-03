package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	entity "github.com/sun053123/go-hexagonal-sqlx/entities"
	handler "github.com/sun053123/go-hexagonal-sqlx/handlers"
	"github.com/sun053123/go-hexagonal-sqlx/logs"
	service "github.com/sun053123/go-hexagonal-sqlx/services"
)

var db *sqlx.DB
var err error

func main() {
	SERVERPORT := os.Getenv("SERVERPORT")

	initDB()
	defer db.Close()

	userEntity := entity.NewUserEntityDB(db)
	// userEntity := entity.NewUserEntityMock()

	userService := service.NewUserService(userEntity)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()

	router.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	router.HandleFunc("/users/{userID:[0-9]+}", userHandler.GetUser).Methods("GET")

	logs.Info("server run at port " + (SERVERPORT))
	err = http.ListenAndServe(SERVERPORT, router)
	if err != nil {
		panic(err)
	}
}

func initDB() {

	HOST := os.Getenv("HOST")
	DBPORT := os.Getenv("DBPORT")
	USER := os.Getenv("USER")
	DBNAME := os.Getenv("DBNAME")
	PASSWORD := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, DBPORT, USER, PASSWORD, DBNAME)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}
