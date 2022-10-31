package main

import (
	"fmt"
)

func main() {
	Player_Array := [3][3]string{
		{"x", "-", "x"},
		{"x", "x", "x"},
		{"x", "x", "x"},
	}
	fmt.Println(Player_Array)

	Player_Array = [3][3]string{
		{"x", "x", "x"},
		{"x", "x", "x"},
		{"x", "x", "x"},
	}
	fmt.Println(Player_Array)

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}
	i := 58
	n := "Frank"
	k := "Happy birthday " + string(n) + "! " + "You are now" + string(i) + "years old!"
	fmt.Println(k)
	name := "Christiane"
	table := 27
	neighbor := "Frank"
	direction := "on the left"
	distance := 23.7834298
	line1 := "Welcome to my party, " + name + "!"
	line2 := fmt.Sprintf("You have been assigned to table %03d. Your table is "+direction+", exactly %.1f meters from here.", table, distance)
	line3 := "You will be sitting next to " + neighbor + "."
	all_together := line1 + "\n" + line2 + "\n" + line3
	fmt.Println(all_together)
}
