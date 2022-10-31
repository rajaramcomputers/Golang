package main

import (
	"fmt"
)

func main() {

	fmt.Println(InterestRate(200.75))
	fmt.Println(Interest(200.75))
	fmt.Println(AnnualBalanceUpdate(200.75))

	balance := 200.75
	targetBalance := 214.88
	fmt.Println(YearsBeforeDesiredBalance(balance, targetBalance))
}

func InterestRate(balance float64) float32 {
	var Interest float32 = 0.00

	if balance <= 0 {
		Interest = 3.213
	} else if balance > 0 && balance < 1000 {
		Interest = 0.5
	} else if balance >= 1000 && balance < 5000 {
		Interest = 1.621
	} else {
		Interest = 2.475
	}

	return Interest
}

func Interest(balance float64) float64 {
	var calculateInterest float64
	calculateInterest = balance * (float64)(InterestRate(balance)/100)
	return calculateInterest
}

func AnnualBalanceUpdate(balance float64) float64 {
	var balanceUpdate float64
	balanceUpdate = balance + Interest(balance)

	return balanceUpdate
}

func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	noofYears := 0

	for balance <= targetBalance {
		noofYears++
		balance = balance + Interest(balance)
	}

	return noofYears
}
