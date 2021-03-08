package main

import "fmt"

func main() {
	dict := make(map[string]int)
	fmt.Printf("\n this is the length of the dict: %v\n", len(dict))
	dict["testing"] = 10
	fmt.Printf("\n this is the length of the dict: %v\n", len(dict))
	delete(dict, "testing")
	fmt.Printf("\n this is the length of the dict: %v\n", len(dict))
}


 