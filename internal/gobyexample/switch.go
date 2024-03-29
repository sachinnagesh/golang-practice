package gobyexample

import (
	"fmt"
	"time"
)

func SwitchExample() {
	//basic switch
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	/*
		commas to separate multiple expressions in the
		same case statement
	*/
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	/*
		switch without an expression is an alternate way
		to express if/else logic. Here we also show how
		the case expressions can be non-constants.
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's afternoon")

	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am an bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know type %T\n", t)

		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
