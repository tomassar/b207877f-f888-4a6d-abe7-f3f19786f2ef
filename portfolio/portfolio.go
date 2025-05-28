package portfolio

type Portfolio struct {
	Stocks     map[string]*Stock
	Allocation map[string]float64
}

type Stock struct {
	Ticker string
	Shares float64

	CurrentPrice func() float64
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
