package main

import "fmt"

// Main function
func main() {

	// Creating and initializing a map
	m := map[int]string{

		90: "hello",
		91: "how",
		92: "are",
		93: "you",
		94: "friend",
	}

	// Iterating map using for rang loop
	for id, pet := range m {

		fmt.Println(id, pet)
	}
}
