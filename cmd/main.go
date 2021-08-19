package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/MaJloe3Jlo/coinconverter/domain"
	"github.com/MaJloe3Jlo/coinconverter/infrasctructure/coinmarket"
	"github.com/MaJloe3Jlo/coinconverter/infrasctructure/convert"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		log.Fatal("Usage: amount(float) currency currency_to_convert. Example: 100.51 USD BTC")
	}

	amountString := arguments[0]
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil {
		log.Fatal("Amount must be in float type")
	}

	if amount <= 0 {
		log.Fatal("Amount must be greater then 0")
	}

	currency := arguments[1]
	currencyToConvert := arguments[2]

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	coinMarket := coinmarket.New(httpClient)

	converter := convert.New(coinMarket)

	result, err := converter.Convert(context.Background(), domain.InputData{Amount: amountString, Currency: currency, CurrencyToConvert: currencyToConvert})
	if err != nil {
		log.Fatalf("Convertation failed. Reason: %v", err)
	}

	fmt.Println("Course is: ", result)
}
