package service

// port

//จะไม่ดึง struct ของ port db มาใช้เนื่องจากบางอันเราก็ไม่อยากให้ออกไป ดังนั้นต้องสร้างมารับเอง
// สร้าง struct เพื่อกำหนด json (data transfer object) DTO
// จะ return อะไรไป client
type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserService interface {
	FindUsers() ([]UserResponse, error)
	FindSingleUser(int) (*UserResponse, error)
}
