package main

import "fmt"

func getLogisticPricing(logisticCode string) (iLogisticPricing, error) {
	switch logisticCode {
	case "sicepat":
		return getPicingSicepat(), nil
	case "jne":
		return getPicingJne(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
