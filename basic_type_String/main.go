package main

import (
	"fmt"
	"strings"
)

func main() {

	fistname, Lastname := "Jeevanantham", "Rangasamy"

	fullName := fistname + " " + Lastname

	fmt.Println(strings.ToUpper(fullName))

}
