package main

import (
	"fmt"
)

func main() {
	sicepat, _ := getLogisticPricing("sicepat")
	printDetails(sicepat)

	jne, _ := getLogisticPricing("jne")
	printDetails(jne)
}

func printDetails(g iLogisticPricing) {
	fmt.Printf("courier: %s", g.getName())
	fmt.Println()
	fmt.Printf("Pricing: %d", g.getPricing())
	fmt.Println()
}
