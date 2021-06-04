package models

import (
	"fmt"
)

// User struct
type User struct {
	ID       int    `json:"id" example:"1" format:"int64"`
	Name     string `json:"name" example:"user name"`
	Lastname string `json:"lastname" example:"user lastname"`
	Email    string `json:"email" example:"user email"`
}

func GetUser(id int) (User, error) {

	if id < 1 {
		return User{}, fmt.Errorf("User %d is not found.", id)
	}

	u := User{
		ID:       id,
		Name:     "Luiz Filipe",
		Lastname: "Miranda de Oliveira",
		Email:    "folivers@gmail.com",
	}

	return u, nil
}
