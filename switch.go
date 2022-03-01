package main

import (
	"fmt"
	"time"
)

func main() {

	i := 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend")
	default:
		fmt.Println("it's the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("I am a bool")
	default:
		fmt.Println("i am an int")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int64:
			fmt.Println("i;m an int64")
		default:
			fmt.Println("Don't Know the type", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
