package main

import (
	"fmt"
)

func equalizeExpenses(expenses map[string]float64) map[string]float64 {
	totalExpenses := 0.0
	for _, amount := range expenses {
		totalExpenses += amount
	}

	averageExpense := totalExpenses / float64(len(expenses))

	balance := make(map[string]float64)
	for person, amount := range expenses {
		balance[person] = amount - averageExpense
	}

	return balance
}

func main() {
	expenses := map[string]float64{
		"John":   100.0,
		"Alice":  150.0,
		"Bob":    120.0,
		"Emily":  90.0,
		"George": 80.0,
	}

	balance := equalizeExpenses(expenses)

	for person, amount := range balance {
		fmt.Printf("%s should %s $%.2f\n", person, getAction(amount), amount)
	}
}

func getAction(amount float64) string {
	if amount > 0 {
		return "receive"
	} else if amount < 0 {
		return "pay"
	}
	return "not owe"
}
