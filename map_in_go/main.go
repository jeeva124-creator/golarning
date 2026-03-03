package main

import "fmt"

func main() {

	marks := map[string]int{
		"Jeeva":  90,
		"barath": 98,
		"naveen": 89,
	}

	fmt.Println(marks["Jeeva"])

	//value and ok pattern
	value, ok := marks["Jeeva"]

	if ok {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key not found")
	}

	for Key, value := range marks {
		fmt.Println(Key, value)
	}

}
