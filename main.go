package main

import (
	"fmt"
)

// Stock represents a stock with a ticker and quantity owned.
type Stock struct {
	Ticker string
	Shares float64

	CurrentPrice func() float64
}

// Portfolio represents a user's portfolio of stocks and their target allocation.
type Portfolio struct {
	Stocks     map[string]*Stock
	Allocation map[string]float64
}

// RebalanceSuggestion represents what to buy or sell to reach the target allocation.
type RebalanceSugestion struct {
	Ticker string
	Action string
	Shares float64
}

// NewPortfolio creates a new portfolio with actual holdings and allocation goals.
func NewPortfolio(stocks []*Stock, allocation map[string]float64) *Portfolio {
	stockMap := make(map[string]*Stock)
	for _, stock := range stocks {
		stockMap[stock.Ticker] = stock
	}
	return &Portfolio{
		Stocks:     stockMap,
		Allocation: allocation,
	}
}

// TotalValue calculates the total market value of the portfolio.
func (p *Portfolio) TotalValue() float64 {
	total := 0.0
	for _, stock := range p.Stocks {
		total += stock.Shares * stock.CurrentPrice()
	}
	return total
}

// Rebalance calculates the buy/sell actions needed to match the target allocation.
func (p *Portfolio) Rebalance() []RebalanceSugestion {
	suggestions := []RebalanceSugestion{}
	totalValue := p.TotalValue()

	// First, calculate current values
	currentValues := map[string]float64{}
	for ticker, stock := range p.Stocks {
		currentValues[ticker] = stock.Shares * stock.CurrentPrice()
	}

	// Then, determine the desired value per ticker based on allocation
	desiredValues := map[string]float64{}
	for ticker, target := range p.Allocation {
		desiredValues[ticker] = target * totalValue
	}

	// Compare current vs desired values and generate suggestions
	for ticker, desiredValue := range desiredValues {
		price := 0.0
		if stock, ok := p.Stocks[ticker]; ok {
			price = stock.CurrentPrice()
		} else {
			// Assume user owns zero if not in current stocks
			price = 100.0 // default price, ideally fetched externally
		}

		currentValue := currentValues[ticker]
		diff := desiredValue - currentValue
		if diff == 0 {
			continue
		}
		action := "buy"
		if diff < 0 {
			action = "sell"
			diff = -diff
		}
		suggestions = append(suggestions, RebalanceSugestion{
			Ticker: ticker,
			Action: action,
			Shares: diff / price,
		})
	}

	return suggestions
}

func main() {
	apple := &Stock{
		Ticker:       "AAPL",
		Shares:       10,
		CurrentPrice: func() float64 { return 150.0 },
	}
	meta := &Stock{
		Ticker:       "META",
		Shares:       5,
		CurrentPrice: func() float64 { return 300.0 },
	}

	allocation := map[string]float64{
		"APL":  0.6,
		"META": 0.4,
	}
	portfolio := NewPortfolio([]*Stock{apple, meta}, allocation)
	suggestions := portfolio.Rebalance()
	for _, s := range suggestions {
		fmt.Printf("%s %.2f sares of %s\n", s.Action, s.Shares, s.Ticker)
	}
}
