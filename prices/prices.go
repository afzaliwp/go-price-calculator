package prices

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/afzaliwp/go-price-calculator/env"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	file, err := os.Open(env.PRICES_FILE)

	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to open resource %s", env.PRICES_FILE)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		priceInFloat, err := strconv.ParseFloat(text, 64)
		if err != nil {
			file.Close()
			fmt.Println(err)
			return errors.New("failed to read the file")
		}

		job.InputPrices = append(job.InputPrices, priceInFloat)
	}

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

	fmt.Println(job.TaxIncludedPrices)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		InputPrices: []float64{},
		TaxRate:     taxRate,
	}
}
