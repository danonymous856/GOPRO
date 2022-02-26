package main

import (
	"fmt"
	"math"
)

const name string = "constant"

func main() {
	fmt.Println(name)

	const n = 5000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
