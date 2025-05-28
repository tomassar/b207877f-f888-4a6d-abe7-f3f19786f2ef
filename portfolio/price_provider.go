package portfolio

type PriceProvider interface {
	GetPrice(ticker string) float64
}

type DefaultPriceProvider struct{}

func (d DefaultPriceProvider) GetPrice(ticker string) float64 {
	return 100.0
}
