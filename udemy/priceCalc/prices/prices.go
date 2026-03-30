package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)


type TaxIncludedPriceJob struct{
	IOManger iomanager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_price"`
	TaxIncludedPrices map[string]string `json:"tax_rate_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	
	lines, err := job.IOManger.ReadLines()

		if err != nil {
			return err
		}

	prices,err := conversion.StringToFloats(lines)

		if err != nil {
			return err
		}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process (doneChan chan bool, errorChan  chan error)  {
	err := job.LoadData()

	// errorChan <- errors.New("Just practice err	")

	if err != nil{
		errorChan <- err
		return
	}
	result := make(map[string]string)
	
		for _, price := range job.InputPrices{
			TaxIncludedPrice := price * (1+job.TaxRate)
			result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
		}
		job.TaxIncludedPrices = result

		job.IOManger.WriteResult(job)
		doneChan <- true 
}

func NewTaxIncludedPriceJob(iom iomanager.IOManager,taxRate float64) (TaxIncludedPriceJob){
	return TaxIncludedPriceJob{
		IOManger: iom,
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}