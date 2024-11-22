package main

import (
	"fmt"
	"github.com/afzaliwp/go-price-calculator/env"
	"github.com/afzaliwp/go-price-calculator/prices"
	"github.com/afzaliwp/go-price-calculator/storage"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.10, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	for index, tax := range taxRates {
		doneChans[index] = make(chan bool)
		fm := storage.NewFileManager(
			env.PRICES_FILE,
			fmt.Sprintf("storage/prices-%.0f-tax.json", tax*100),
		)
		taxIncludedPriceJob := prices.NewTaxIncludedPriceJob(fm, tax)
		go taxIncludedPriceJob.Process(doneChans[index])
	}

	for _, doneChan := range doneChans {
		<-doneChan
	}
}
