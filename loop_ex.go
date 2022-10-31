package main

import "fmt"

func main() {
	/* //breakpoint:
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x == y && y == 1 {
				//break breakpoint
			}
			fmt.Println("x,y", x, y)
		}
	} */
	card1 := "ten"
	card2 := "six"
	dealer := "six"

	result := FirstTurn(card1, card2, dealer)
	fmt.Println("the result is", result)
}
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}

}
func FirstTurn(card1, card2, dealerCard string) string {
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case ParseCard(card1)+ParseCard(card2) == 21:
		if dealerCard != "ten" && dealerCard != "ace" && dealerCard != "jack" && dealerCard != "queen" && dealerCard != "king" {
			return "W"
		} else {
			return "S"
		}
	case ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20:
		return "S"
	case ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H"
		} else {
			return "S"
		}
	case ParseCard(card1)+ParseCard(card2) <= 11:
		return "H"
	}
	return "H"
}
