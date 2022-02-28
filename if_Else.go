package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("odd number\n")
	} else {
		fmt.Println("even number\n")
	}

	if 8%2 == 0 {
		x := "even number"
		fmt.Println(x)
	}
}
