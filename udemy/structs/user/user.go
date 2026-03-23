package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName string
	birthDate string
	createdAt time.Time
}

type Admin struct {
	email string
	password string
	User
}

// func outputUserDetails (u *User) {
	// 	fmt.Println((*u).firstName, u.lastName, u.birthDate)
	// }

func (u *User) OutputUserDetails () {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName () {
	(*u).firstName = ""
	u.lastName = ""
}
func NewAdmin(email, password string) *Admin {
	return &Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "_ADMIN_",
			birthDate: "---",
			createdAt: time.Now(),
		},
	}
}

func New (firstName,lastName, birthDate string) (*User, error){
	if firstName == "" || lastName == "" || birthDate == ""{
		return nil, errors.New("First name, Last name, birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}
