package entity

//Port ที่เป็น plug ว่าจะกำหนดการทำงาน

type User struct {
	UserID    int    `db:"id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	Username  string `db:"username"`
	Email     string `db:"email"`
	Password  string `db:"password"`
}

type UserEntity interface {
	GetAll() ([]User, error)
	GetById(int) (*User, error)
}

// GetById ไม่สามารถ return nil ได้เพราะต้องการสักตัวนึง
