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

	// if with short statement

	item := 3
	priseperitem := 49

	if total := item * priseperitem; total > 100 {
		fmt.Println("your eligable for Shipping ")
	} else {
		fmt.Println("Not elgibale")
	}
}
