package main

type sicepat struct {
	logisticPricing
}

func getPicingSicepat() iLogisticPricing {
	return &sicepat{logisticPricing{pricing: 100, name: "sicepat"}}
}
