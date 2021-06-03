package models

// User struct
type User struct {
	ID       int    `json:"id" example:"1" format:"int64"`
	Name     string `json:"name" example:"user name"`
	Lastname string `json:"lastname" example:"user lastname"`
	Email    string `json:"email" example:"user email"`
}
