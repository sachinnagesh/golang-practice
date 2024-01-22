package gobyexample

import "fmt"

func IfElseExample() {
	//if without else
	if 4%2 == 0 {
		fmt.Println("4 is even")
	}

	//if with else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//if with logical operator
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 is even")
	}

	/**
	A statement can precede conditionals;
	any variables declared in this statement
	are available in the current and all subsequent branches
	**/

	if n := 9; n < 0 {
		fmt.Println(n, " is negative")
	} else if n > 10 {
		fmt.Println(n, " is greater then 10")
	} else {
		fmt.Println(n, " is positive and less than 10")
	}

}
