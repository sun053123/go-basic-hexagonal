package entity

import "errors"

type userEntityMock struct {
	users []User
}

func NewUserEntityMock() UserEntity {
	users := []User{
		{UserID: 777, Username: "test1", Email: "test@test.com", Password: "123456", CreatedAt: "idk", UpdatedAt: "idl"},
		{UserID: 888, Username: "test2", Email: "test@test.com", Password: "123456", CreatedAt: "idk", UpdatedAt: "idl"},
		{UserID: 999, Username: "test3", Email: "test@test.com", Password: "123456", CreatedAt: "idk", UpdatedAt: "idl"},
	}

	return userEntityMock{users: users}
}

func (ent userEntityMock) GetAll() ([]User, error) {
	return ent.users, nil
}

func (ent userEntityMock) GetById(id int) (*User, error) {
	for _, user := range ent.users {
		if user.UserID == id {
			return &user, nil
		}

	}
	return nil, errors.New("TEST user not found")
}
