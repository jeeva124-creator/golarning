package main

import "fmt"

func main() {
	isLogged := true
	isAdmin := false
	hasScbscription := true
	// And &&
	conOpenDashbord := isLogged && hasScbscription

	canDelitePost := isAdmin || (isLogged && hasScbscription)
	fmt.Println(canDelitePost, conOpenDashbord)

}
