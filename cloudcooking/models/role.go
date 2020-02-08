package models

//Role defines user roles
type Role struct {
	BaseModel
	Name string
	Users []User
}
