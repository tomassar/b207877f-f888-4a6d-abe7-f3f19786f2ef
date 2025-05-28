package main

import (
	"fmt"

	"github.com/tomassar/b207877f-f888-4a6d-abe7-f3f19786f2ef/portfolio"
)

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
	portfolio := portfolio.NewPortfolio([]*portfolio.Stock{apple, meta}, allocation)
	suggestions := portfolio.Rebalance()
	for _, s := range suggestions {
		fmt.Printf("%s %.2f shares of %s\n", s.Action, s.Shares, s.Ticker)
	}
}
