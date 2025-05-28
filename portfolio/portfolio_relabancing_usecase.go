package portfolio

type RebalanceSuggestion struct {
	Ticker string
	Action string
	Shares float64
}

// Rebalance calculates the buy/sell actions needed to match the target allocation.
func (p *Portfolio) Rebalance() []RebalanceSuggestion {
	suggestions := []RebalanceSuggestion{}
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
		suggestions = append(suggestions, RebalanceSuggestion{
			Ticker: ticker,
			Action: action,
			Shares: diff / price,
		})
	}

	return suggestions
}
