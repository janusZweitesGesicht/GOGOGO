package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)


const accountBalanceFile = "balance.txt"

func writeBalanceToFile (balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func getBalanceFromFile () (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	
	if err != nil {
		return 1000, errors.New("Failed to parse balance file")
	}

	return balance, nil
}



func main () {
	var accountBalance, err = getBalanceFromFile()
	if err != nil{
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------------")
		panic ("Cant continue sry")
	}
	fmt.Println("Welcome to Go Bank!")

	for {

	fmt.Println("What do u want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Ur choise:")
	fmt.Scan(&choice)
	fmt.Println("Your choise is:", choice)

	//wantTodoSmth :=choise = 1 // == bool

	switch choice {
		case 1: fmt.Println("Your balance is:", accountBalance)
		case 2:fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0 {
			fmt.Println("Deposit amount must be > 0")
			return
		}
		
		accountBalance += depositAmount
		fmt.Println("Balance updated: ", accountBalance)
		writeBalanceToFile(accountBalance)
	case 3:fmt.Print("Withdrawal amount : ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		if withdrawalAmount <= 0 {
			fmt.Println("Deposit amount must be > 0")
			continue
		}
		if withdrawalAmount > accountBalance{
			fmt.Println("u CANT WITHDROW MORE THAN U HAVE")
			continue
		}
		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated: ", accountBalance)
		writeBalanceToFile(accountBalance)

	default: 		
		fmt.Print("Goodbay")
		fmt.Println("Thanks for using our bank")
		return
	}
	}
}

// Welcome to Go Bank!
// What do you want to do?
// 1. Check balance
// 2. Deposit money
// 3. Withdraw money
// 4. Exit
// Your choice: 