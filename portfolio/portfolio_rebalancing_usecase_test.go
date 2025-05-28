package portfolio

import (
	"testing"
)

func TestPortfolioRebalancing(t *testing.T) {
	tests := []struct {
		name        string
		stocks      []*Stock
		allocation  map[string]float64
		expected    map[string]RebalanceSuggestion
		expectError bool
	}{
		{
			name: "basic rebalancing - buy and sell",
			stocks: []*Stock{
				{Ticker: "AAPL", Shares: 10, CurrentPrice: func() float64 { return 150.0 }},
				{Ticker: "META", Shares: 5, CurrentPrice: func() float64 { return 300.0 }},
			},
			allocation: map[string]float64{
				"AAPL": 0.6,
				"META": 0.4,
			},
			expected: map[string]RebalanceSuggestion{
				"META": {Ticker: "META", Action: "sell", Shares: 1.0},
				"AAPL": {Ticker: "AAPL", Action: "buy", Shares: 2.0},
			},
			expectError: false,
		},
		{
			name: "already balanced portfolio",
			stocks: []*Stock{
				{Ticker: "AAPL", Shares: 6, CurrentPrice: func() float64 { return 100.0 }},
				{Ticker: "META", Shares: 4, CurrentPrice: func() float64 { return 100.0 }},
			},
			allocation: map[string]float64{
				"AAPL": 0.6,
				"META": 0.4,
			},
			expected:    map[string]RebalanceSuggestion{},
			expectError: false,
		},
		{
			name: "invalid allocation - doesn't sum to 1.0",
			stocks: []*Stock{
				{Ticker: "AAPL", Shares: 10, CurrentPrice: func() float64 { return 150.0 }},
			},
			allocation: map[string]float64{
				"AAPL": 0.8,
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			portfolio, err := NewPortfolio(tt.stocks, tt.allocation)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			suggestions := portfolio.Rebalance()

			if len(suggestions) != len(tt.expected) {
				t.Errorf("expected %d suggestions, got %d", len(tt.expected), len(suggestions))
				return
			}

			actualMap := make(map[string]RebalanceSuggestion)
			for _, suggestion := range suggestions {
				actualMap[suggestion.Ticker] = suggestion
			}

			for ticker, expected := range tt.expected {
				actual, exists := actualMap[ticker]
				if !exists {
					t.Errorf("missing suggestion for ticker %s", ticker)
					continue
				}
				if actual.Action != expected.Action {
					t.Errorf("ticker %s: expected action %s, got %s", ticker, expected.Action, actual.Action)
				}
				if abs(actual.Shares-expected.Shares) > 0.01 {
					t.Errorf("ticker %s: expected shares %.2f, got %.2f", ticker, expected.Shares, actual.Shares)
				}
			}
		})
	}
}

func TestPortfolioTotalValue(t *testing.T) {
	stocks := []*Stock{
		{Ticker: "AAPL", Shares: 10, CurrentPrice: func() float64 { return 150.0 }},
		{Ticker: "META", Shares: 5, CurrentPrice: func() float64 { return 300.0 }},
	}
	allocation := map[string]float64{
		"AAPL": 0.6,
		"META": 0.4,
	}

	portfolio, err := NewPortfolio(stocks, allocation)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTotal := 10*150.0 + 5*300.0
	actualTotal := portfolio.TotalValue()

	if actualTotal != expectedTotal {
		t.Errorf("expected total value %.2f, got %.2f", expectedTotal, actualTotal)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
