package main

type logisticPricing struct {
	name    string
	pricing int
}

func (l *logisticPricing) getPricing() int {
	return l.pricing
}

func (l *logisticPricing) getName() string {
	return l.name
}
