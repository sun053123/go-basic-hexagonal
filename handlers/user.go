package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	service "github.com/sun053123/go-hexagonal-sqlx/services"
)

// presentation layer
// มีแต่ adapter ไม่มี port ดังนั้นไม่มี interface

//เป็น adapter จึงไม่ expose และไม่ใช้ entity โดยตรง ใช้แต่ service เท่านั้น
type userHandler struct {
	userServ service.UserService
}

func NewUserHandler(userServ service.UserService) userHandler {
	return userHandler{userServ: userServ}
}

func (handl userHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := handl.userServ.FindUsers()
	// error ตัวนี้ต้องมาจากทาง business แล้ว
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (handl userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// ดึง param  ID มาจาก url
	// ไม่ handle err เพราะประกาศ regex path ไว้แล้ว ไม่มีทางเป็ string ได้
	userID, _ := strconv.Atoi(mux.Vars(r)["userID"])

	user, err := handl.userServ.FindSingleUser(userID)
	if err != nil {

		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
