package gobyexample

import "fmt"

func DeferExample() {
	defer fmt.Println("world")
	test1()
	test2()

	fmt.Println("hello")
}

func test1() {
	fmt.Println("test1 : first")
	defer fmt.Println("test1 : defer")
	fmt.Println("test1 : last")
}

func test2() {
	fmt.Println("test2 : first")
	defer fmt.Println("test2 : defer")
	fmt.Println("test2 : last")
}
