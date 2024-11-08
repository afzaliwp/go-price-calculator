package main

import (
	"github.com/afzaliwp/go-price-calculator/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.10, 0.15}

	for _, tax := range taxRates {
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(tax)
		taxIncludedPriceJob.Process()
	}

}
