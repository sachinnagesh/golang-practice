package gobyexample

import (
	"fmt"
)

const s string = "constant"

func ConstantExample() {

	fmt.Println(s)

	const n = 50000000000

	const d = 3e20 / n
	fmt.Println(d)

}
