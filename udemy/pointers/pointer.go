package main

import "fmt"

func main () {
	age := 43
	agePtr := &age

	fmt.Println("Age: ", *agePtr)

	editToAdultYears(agePtr)
	fmt.Println(age)

}

func editToAdultYears (age *int) {
	// return *age - 18
	*age = *age - 18
}