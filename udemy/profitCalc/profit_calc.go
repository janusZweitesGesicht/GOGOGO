package main

import (
	"errors"
	"fmt"
	"os"
)

const resultCalcFile = "calcResults.txt"

// Goals
// 1) Validate user input
// Show error message & exit if invalid input is provided
// - No negative numbers
// - Not 0
// 2) Store calculated results into file

func main() {
	revenue, err := getUserInput("Revenue: ")
	if (err != nil) {
		panic(err)}
	expenses, err := getUserInput("Expenses: ")
	if (err != nil) {
		panic(err)}
	taxRate, err := getUserInput("Tax Rate: ")
	if (err != nil) {
		panic(err)}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeResToFile(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if(userInput <= 0) {
		return 0, errors.New("Invalid input!!!")
	}
	return userInput, nil
}

func writeResToFile (ebt, profit, ratio float64) {
	textToFile := "ebt: " + fmt.Sprint(ebt) + "\nprofit: " + fmt.Sprint(profit) + "\nratio: " + fmt.Sprint(ratio)
	//results := fmt. Sprintf(" EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(resultCalcFile, []byte(textToFile), 0644)
}