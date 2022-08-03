package roman_number

import "fmt"

func maximumWealth(accounts [][]int) int {
	richest := make(map[int]int)
	for customer := 0; customer < len(accounts); customer++ {
		money := 0
		for bank := 0; bank < len(accounts[customer]); bank++ {
			money += accounts[customer][bank]
		}
		fmt.Println("Customer", customer, " has", money)
		richest[customer] = money
	}

	wealthiest := 0
	for _, money := range richest {
		if money > wealthiest {
			wealthiest = money
		}
	}
	return wealthiest

	return 0
}
