package main

import "fmt"

func main() {
	var number int
	fmt.Scanln(&number)

	if number >= 5 {
		fmt.Print("Excellent!")
	}
}
