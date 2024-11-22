package prices

import (
	"fmt"
	"math"

	"github.com/afzaliwp/go-price-calculator/storage"
)

type TaxIncludedPriceJob struct {
	IOManager         storage.FileManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	data, err := job.IOManager.ReadFile()
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

	err := job.IOManager.OutputJsonFile(job)

	if err != nil {
		fmt.Errorf("error saving prices: %s", err.Error())
		return
	}

	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(fm *storage.FileManager, taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		IOManager:   *fm,
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
