package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Printf("%-20v %-5v %-15v $ %4v\n", "SPACELINE", "Days", "Triptype", "Price")
	fmt.Printf("========================================\n")
	var count = 15
	for count > 0 {

		switch randomNumber := rand.Intn(3); randomNumber {
		case 0:
			fmt.Printf("%-20v", "Space Adventures")
		case 1:
			fmt.Printf("%-20v", "SpaceX")
		case 2:
			fmt.Printf("%-20v", "Virgin Galactic")
		}

		const distance = 62100000
		var speed = rand.Intn(14) + 16
		var price = speed - 16 + 36
		var days = distance / speed / 24 / 60 / 60

		switch randomNumber := rand.Intn(2); randomNumber {
		case 0:
			fmt.Printf("%-5v %-15v $ %4v\n", days, "One-Way", price)
		case 1:
			fmt.Printf("%-5v %-15v $ %4v\n", days*2, "Round-Trip", price*2)
		}

		count--
	}
}
