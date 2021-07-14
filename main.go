package main

import "fmt"

func main() {
	//fmt.Println("Hello World")
	//fmt.Println("I am Isfandiyor Zabirov")

	minAmount := 5_000
	purchase := 15_000
	percentCashback := 15
	limitCashback := 2_000
	cashback := 0
	const fullPercent int = 100

	// 1. Можно ли дать кэшбек
	if purchase >= minAmount {
		cashback = purchase * percentCashback / fullPercent
		fmt.Println("Prediction cashback is", cashback)

	}
	// 2. Превышание кэшбека
	if cashback > limitCashback {
		cashback = limitCashback
	}
	fmt.Println("Cashback is", cashback)
}

//git init
//git add .
//git commit -m "Any message"