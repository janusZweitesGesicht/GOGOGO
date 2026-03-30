package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)



func main(){

	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	for i := range taxRates {
		doneChans[i] = make(chan bool)
	}
	for i, taxRate := range taxRates{
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		// priceJob := prices.NewTaxIncludedPriceJob(cmdm, taxRate)
		// err := priceJob.Process()
		go priceJob.Process(doneChans[i])


		// if err != nil {
		// 	fmt.Println(err)
		// }
	}
	for _, doneChan := range doneChans {
		<-doneChan
	}
}