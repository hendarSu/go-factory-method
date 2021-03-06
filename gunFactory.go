package main

import "fmt"

func getGun(gunType string) (iGun, error) {
	switch gunType {
	case "ak47":
		return newAk47(), nil
	case "handgun":
		return newHandGun(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}
