package main

import (
	"fmt"
	"log"

	"github.com/tomassar/b207877f-f888-4a6d-abe7-f3f19786f2ef/portfolio"
)

type SimplePriceProvider struct{}

func (s SimplePriceProvider) GetPrice(ticker string) float64 {
	prices := map[string]float64{
		"AAPL":  150.0,
		"META":  300.0,
		"GOOGL": 120.0,
	}
	if price, ok := prices[ticker]; ok {
		return price
	}
	return 100.0 // fallback
}

func main() {
	apple := &portfolio.Stock{
		Ticker:       "AAPL",
		Shares:       10,
		CurrentPrice: func() float64 { return 150.0 },
	}
	meta := &portfolio.Stock{
		Ticker:       "META",
		Shares:       5,
		CurrentPrice: func() float64 { return 300.0 },
	}

	allocation := map[string]float64{
		"AAPL": 0.6,
		"META": 0.4,
	}
	portfolio, err := portfolio.NewPortfolio([]*portfolio.Stock{apple, meta}, allocation)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally set a custom price provider
	portfolio.SetPriceProvider(SimplePriceProvider{})

	suggestions := portfolio.Rebalance()
	for _, s := range suggestions {
		fmt.Printf("%s %.2f shares of %s\n", s.Action, s.Shares, s.Ticker)
	}
}
