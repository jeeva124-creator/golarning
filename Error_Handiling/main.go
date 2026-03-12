package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {

	//GO dont use expations for normal failures
	//   function  =return errors as normal return value
   if  err :
}

func run() error {
	intput := "3"
	level, err := parsLevel(intput)
	if err != nil {
		return err
	}
	fmt.Println("selected level", level)
	return nil
}
func parsLevel(s string) (int, error) {
	n, err := strconv.Atoi(s)

	if err != nil {
		return 0, fmt.Errorf("level mustbe number")
	}
	if n < 1 || n > 5 {
		return 0, fmt.Errorf("Level must be a 1 and 5")
	}
	return n, nil
}
