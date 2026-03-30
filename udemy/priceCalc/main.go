package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)



func main(){

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))
	
	for i, taxRate := range taxRates{
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
	
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		// err := priceJob.Process()
		go priceJob.Process(doneChans[i], errorChans[i])


		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

	for i := range taxRates {
		select {
		case err:= <- errorChans[i]:
			if err != nil{
				fmt.Println(err)
			}
		case <- doneChans[i]:
		fmt.Println("Done!")	
		} 
	}
}