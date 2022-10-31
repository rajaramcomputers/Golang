package main

import "fmt"

func main() {
	m := map[string]bool{
		"x": true,
		"y": false,
	}
	fmt.Println(m)
	players := make(map[string]int)

	// add value to the map
	players["ram"] = 100
	players["vardhini"] = 90
	players["meera"] = 100
	fmt.Println(players)

	fmt.Println(players["ram"], "next", players["meera"])

	//deleting a map value

	delete(players, "ram")
	fmt.Println("Players after Deletion", players)

	//check whether the key exists and get values
	//values will be the default value if the key does not exist

	v, ok := players["meera"]
	fmt.Println("key value", v, "Ok status", ok)

	v, ok = players["ram"]
	fmt.Println("key value", v, "Ok Status", ok)

	//map of slices

	mp := map[string][]string{
		"admins": []string{"turkia", "haifa"},
		"users":  []string{"hamid", "senthil"},
	}
	fmt.Println("Map of Slices", mp)

}
