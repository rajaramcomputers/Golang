package main

import "fmt"

func main() {

	//range usage in the array
	arr := [4]string{"ram", "krishna", "govinda", "Achuda"}

	for key, value := range arr {
		fmt.Println("range of array key", key, "value", value)
	}
	// range usage in the slice
	slc := []string{"x", "y", "z"}
	for k, v := range slc {
		fmt.Println("range of the slice key", k, "value", v)
	}
	//range usage in the map

	mp := map[string]string{
		"admin": "ranjith",
		"user":  "senthil",
	}

	for map_key, map_value := range mp {
		fmt.Println("Key of the Map", map_key, "Value", map_value)
	}

}
