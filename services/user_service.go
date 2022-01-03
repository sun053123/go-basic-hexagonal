package service

import (
	"database/sql"

	entity "github.com/sun053123/go-hexagonal-sqlx/entities"
	"github.com/sun053123/go-hexagonal-sqlx/errs"
	"github.com/sun053123/go-hexagonal-sqlx/logs"
)

//  Business Logic ใน Logic มีการ implement สองฝั่ง

// Adapter

// สร้าง struct ไว้เก็บ data  มาใช้งานเพราะมันไม่มี data ใช้เอง รับมาจาก ส่วนอื่น
// อ้างการเข้าถึงผ่าน interface

type userService struct {
	userEnt entity.UserEntity //ผ่าน interface เพราะจะไม่มีการไปยุ่งใดๆ กับ db เลย จะต้องเรียกผ่าน port เท่านั้น
}

//อย่างแรกต้อง instant มันขึ้นมาก่อนโดยส่่งมาจะใช้ ตาม entity interface อีกที
func NewUserService(userEnt entity.UserEntity) UserService {
	return userService{userEnt: userEnt}
}

func (serv userService) FindUsers() ([]UserResponse, error) {

	users, err := serv.userEnt.GetAll() // เรียกใช้งานฝั่งฟังก์ชันอง db เพราะตรงนั้นมีหน้าที่ตามการทำงาน ดึงข้อมูล
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	// จากตอนนี้ได้ข้อมูลการ getall users มาแล้วจากฝั่ง DB แต่จะได้ ข้อมูล ทั้ง struct ของฝั่ง DB มาเลย
	// เราต้องการไม่ให้ client รู้ข้อมูลทั้งหมดเราจะต้องทำตาม struct ของฝั่ง service ที่เราสร้างไว้

	userResps := []UserResponse{}
	//loop apend ให้มีข้อมูลไปตามที่กำหนด ข้อสังเกตุคือ user กับ userResps ข้อมูลคนละตัว จะต้อง bind ให้ตรงกันก่อน append
	// คือ ปั้น object ใหม่
	for _, item := range users {
		userresponse := UserResponse{
			Username: item.Username,
			Email:    item.Email,
		}
		// ปั้น obj userResp เพื่อใส่ใน slide userResps
		userResps = append(userResps, userresponse)
	}
	return userResps, nil
}

func (serv userService) FindSingleUser(id int) (*UserResponse, error) {
	user, err := serv.userEnt.GetById(id)
	if err != nil {
		//hadle err ที่หาไม่เจอ เพราะไม่ตรงกับ db ปั้นให้ส่งแค่ว่า  ข้อมูล nil และ new err "หาไม่เจอ"
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found !")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	userResp := UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}

	return &userResp, nil

}
