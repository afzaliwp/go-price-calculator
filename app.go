package main

import (
	"fmt"
	"github.com/afzaliwp/go-price-calculator/env"
	"github.com/afzaliwp/go-price-calculator/prices"
	"github.com/afzaliwp/go-price-calculator/storage"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.10, 0.15}

	for _, tax := range taxRates {
		fm := storage.NewFileManager(
			env.PRICES_FILE,
			fmt.Sprintf("storage/prices-%.0f-tax.json", tax*100),
		)
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		taxIncludedPriceJob.Process()
	}

}
