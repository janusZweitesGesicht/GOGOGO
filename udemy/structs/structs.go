package main

import (
	"fmt"
	"example.com/structs/structs/user"
)

func main() {
	userfirstName := getUserData("Please enter your first name: ")
	userlastName := getUserData("Please enter your last name: ")
	userbirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User
	
	appUser, err := user.New(userfirstName, userlastName, userbirthDate)
	if err != nil {
		panic(err)
	}



	admin := user.NewAdmin("test@example.com", "1234")
	admin.User.OutputUserDetails()
	admin.ClearUserName()
	admin.OutputUserDetails()
	//appUser = User{
	// 	userfirstName,
	// 	userlastName,
	// 	userbirthdate,
	// 	time.Now(),
	// }
	// ... do something awesome with that gathered data!
	// outputUserDetails(&appUser)	
	// fmt.Println(firstName, lastName, birthDate)
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
