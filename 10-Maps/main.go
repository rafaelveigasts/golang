package main

import "fmt"

func main() {
	// Create a map
	// var m map[string]int
	// m = make(map[string]int)
	m := make(map[string]int)

	// Add key-value pairs to the map
	m["k1"] = 7
	m["k2"] = 13

	// Print the map
	fmt.Println("map:", m)
	fmt.Println("map:", m["k1"])

	// Get a value for a key with the name of the key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// Get a value for a key with the name of the key
	// fmt.Println("len:", len(m))

	// Remove a key-value pair from the map
	// delete(m, "k2")
	// fmt.Println("map:", m)

	// Check if a key is present in the map
	// _, prs := m["k2"]
	// fmt.Println("prs:", prs)

	// Declare and initialize a new map in a single line
	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map:", n)
	// fmt.Println("map:", n["foo"])
}