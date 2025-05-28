package portfolio

import "fmt"

type Portfolio struct {
	Stocks        map[string]*Stock
	Allocation    map[string]float64
	PriceProvider PriceProvider
}

type Stock struct {
	Ticker string
	Shares float64

	CurrentPrice func() float64
}

// NewPortfolio creates a new portfolio with actual holdings and allocation goals.
func NewPortfolio(stocks []*Stock, allocation map[string]float64) (*Portfolio, error) {
	// Validate that allocations sum to 1.0
	total := 0.0
	for _, alloc := range allocation {
		total += alloc
	}
	if total < 0.99 || total > 1.01 { // Allow small floating point tolerance
		return nil, fmt.Errorf("allocations must sum to 1.0, got %.3f", total)
	}

	stockMap := make(map[string]*Stock)
	for _, stock := range stocks {
		stockMap[stock.Ticker] = stock
	}
	return &Portfolio{
		Stocks:        stockMap,
		Allocation:    allocation,
		PriceProvider: DefaultPriceProvider{},
	}, nil
}

// TotalValue calculates the total market value of the portfolio.
func (p *Portfolio) TotalValue() float64 {
	total := 0.0
	for _, stock := range p.Stocks {
		total += stock.Shares * stock.CurrentPrice()
	}
	return total
}

// SetPriceProvider allows setting a custom price provider.
func (p *Portfolio) SetPriceProvider(provider PriceProvider) {
	p.PriceProvider = provider
}
