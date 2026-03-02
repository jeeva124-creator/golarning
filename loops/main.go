package main

import "fmt"

func main() {

	start := 0
	end := 1
	n := 5
	for i := start; i <= n; i++ {
		temp := start + end
		end = start
		start = temp
	}

	fmt.Println(start)
	// 0112358
}
