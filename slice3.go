package main

import "fmt"

func main() {

	dest := make([]string, 3)
	fmt.Println(dest)

	var src = []string{"ram", "krishna", "govinda"}

	//copy

	copy(dest, src)
	fmt.Println(dest)

	src[0] = "abi"

	fmt.Println(dest)
	fmt.Println(src)

	//when you reference the scenario is different

	newsrc := src[0:3]
	fmt.Println(newsrc)

	src[0] = "vardhini"
	fmt.Println("Src: ", src)
	fmt.Println("Dest: ", dest)
	fmt.Println("New Source: ", newsrc)

	//some more functionality with the slice

	fmt.Println(src)
	fmt.Println("src :3", src[:3])
	fmt.Println("src 0:2", src[0:2])
	fmt.Println("src 2:", src[2:])

	gameMap := [][]string{
		make([]string, 3),
		make([]string, 5),
		([]string{"-", "-"}),
	}
	fmt.Println(gameMap)
	//make change in the size

	gameMap[0] = make([]string, 6)
	gameMap[1] = make([]string, 10)

	fmt.Println(gameMap)

}
