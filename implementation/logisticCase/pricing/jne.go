package main

type jne struct {
	logisticPricing
}

func getPicingJne() iLogisticPricing {
	return &sicepat{logisticPricing{pricing: 150, name: "jne"}}
}
