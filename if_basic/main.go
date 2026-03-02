package main

import "fmt"

func main() {
	score := 70

	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else if score > 50 {
		fmt.Println("c")
	}
}
