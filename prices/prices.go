package prices

import (
	"fmt"
	"math"

	"github.com/afzaliwp/go-price-calculator/env"
	"github.com/afzaliwp/go-price-calculator/storage"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	data, err := storage.ReadFile(env.PRICES_FILE)
	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = data

	return nil
}

func (job TaxIncludedPriceJob) Process() {
	job.LoadPrices()
	job.TaxIncludedPrices = make(map[string]float64, len(job.InputPrices))

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)             //Calculated the taxed price
		TaxIncludedPrice = math.Round(TaxIncludedPrice*100) / 100 //Round to two decimals
		job.TaxIncludedPrices[fmt.Sprintf("%.2f", price)] = TaxIncludedPrice
	}

	err := storage.SaveJson(
		fmt.Sprintf("storage/prices-%.0f-tax.json", job.TaxRate*100),
		job.TaxIncludedPrices,
	)

	if err != nil {
		fmt.Errorf("error saving prices: %s", err.Error())
		return
	}

	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
