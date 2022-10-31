package main

import (
	"fmt"
)

func main() {
	var empty []int
	withData := []int{0, 1, 2, 3, 4, 5}

	fmt.Println(empty)
	fmt.Println(withData)

	withData[1] = -1

	fmt.Println(withData)

	getaValue := withData[1]

	fmt.Println(getaValue)

	newSlice := withData[2:4]
	fmt.Println(newSlice)

	newSlice = withData[:2]
	fmt.Println(newSlice)

	withData = append(withData, 6, 7)

	fmt.Println(withData)

	find(10, 10, 20, 30, 40)

	cards := FavoriteCards()
	fmt.Println(cards)

	card := GetItem([]int{1, 2, 4, 1}, 2)

	fmt.Println(card)

	card = GetItem([]int{1, 2, 4, 1}, 10) // card == -1
	fmt.Println(card)

	/* index := 2
	newCard := 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards) */

	slice := []int{3, 2, 6, 4, 8}
	cards = PrependItems(slice)
	fmt.Println(cards)
	// Output: [3 2 6 4 8]

	/* slice := []int{3, 2, 6, 4, 8}
	cards = PrependItems(slice, 5, 1)
	fmt.Println(cards)
	// Output: [5 1 3 2 6 4 8] */

	index := 2
	newCard := 6
	cards = SetItem([]int{1, 2, 4, 1}, index, newCard)
	fmt.Println(cards)

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 2)
	fmt.Println("The Removed Items", cards)

	cards = RemoveItem([]int{3, 2, 6, 4, 8}, 11)
	fmt.Println(cards)

	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDay))

	birdsPerDay1 := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(BirdsInWeek(birdsPerDay1, 2))

	birdsPerDay2 := []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDay2))
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			return
		}
	}

	fmt.Println(num, "not found in ", nums)
}

func FavoriteCards() []int {
	cards := []int{2, 6, 9}
	return cards
}

func GetItem(slice []int, index int) int {

	if index <= len(slice) {
		return slice[index]
	} else {
		return -1
	}
}
func TotalBirdCount(birdsPerDay []int) int {
	var sum = 0
	for i := 0; i < len(birdsPerDay); i++ {
		sum = sum + birdsPerDay[i]

	}
	return sum
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	var sum = 0
	var startpoint int
	if week == 1 {
		startpoint = 0
	} else {
		startpoint = (week - 1) * 7
	}
	var endpoint = startpoint + 6

	for j := startpoint; j <= endpoint; j++ {
		sum = sum + birdsPerDay[j]
	}

	return sum
}

func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i = i + 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
		//newslice = append(newslice, birdsPerDay[i])
	}
	return birdsPerDay
}

func SetItem(slice []int, index, value int) []int {
	if index > 0 && index <= len(slice) {
		slice[index] = value
		return slice
	} else {
		slice = append(slice, value)
		return slice

	}
}

func PrependItems(slice []int, values ...int) []int {
	// to create a new slice witht the passed values
	var new_slice []int
	for _, v := range values {

		new_slice = append(new_slice, v)
	}
	//to preappend the values with the new slice created
	for _, v := range slice {

		new_slice = append(new_slice, v)
	}

	return new_slice

}

func RemoveItem(slice []int, index int) []int {
	if index > 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}

}
