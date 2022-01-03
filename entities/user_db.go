package entity

import "github.com/jmoiron/sqlx"

//Adapter ที่เชื่อมต่อไปยัง Database

// package อื่นห้ามอ้างถึง Adapter ต้อง instant ผ่าน func New
// ไม่ให้ใครเข้าถึง database โดยตรง
type userEntityDB struct {
	db *sqlx.DB
}

// func constructer
// ถ้าจะเรียก func new จะต้องส่ง db มาให้ ที่จะเอาไป instant struct
// สามารถทำ database pool ได้ เพราะไม่ได้บังคับให้่ส่ง db ที่ specific มา
func NewUserEntityDB(db *sqlx.DB) UserEntity {
	return userEntityDB{db: db}
}

// func ถ้าจะทำได้ต้อง conform ไปตาม interface ที่ประกาศ function ไว้
func (ent userEntityDB) GetAll() ([]User, error) {
	users := []User{}
	query := "SELECT id, created_at, updated_at, username, email, password FROM users"
	err := ent.db.Select(&users, query)
	if err != nil {
		return nil, err
		//ฝั่ง entity ไม่ได้มีส่วนจัดการ business ดังนั้นจึงโยน error ให้ส่วนอื่นจัดการ
	}
	return users, nil
}

func (ent userEntityDB) GetById(id int) (*User, error) {
	user := User{}
	query := "SELECT id, created_at, updated_at, username, email, password FROM users WHERE id=$1"
	err := ent.db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
