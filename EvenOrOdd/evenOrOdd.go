package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Print("even")
	} else {
		fmt.Print("odd")
	}
}
