package internal

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"strconv"
)

func PrintHelloWorld() {
	fmt.Println("Hello World !!!")
}

func MathRandomExamp() {
	fmt.Println(rand.Int())
}

func PrintfExamp() {
	fmt.Printf("\nSquare root of 49 id %f", math.Sqrt(7))
}

func VariablesExamp() {
	fmt.Println("\n internal : TestVariables")
	var i int = 10
	var j float32 = 10.14
	var k float64 = 44.900033
	var l string = "I am string"
	var m bool = true
	var n complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Println(i, j, k, l, m, n)

	i = 20
	fmt.Println(i)

}

func DefaultValuesExamp() {
	fmt.Println("\n internal : GetDefaultValues ")
	var i int
	var j float32
	var k float64
	var l bool
	var m string
	fmt.Printf("%v , %v , %v, %v, %q \n", i, j, k, l, m)
}

func TypeInferenceExamp() {
	fmt.Println("\ninternal : TypeInferenceExamp")
	i := 99
	fmt.Println(i)
	var j int = 37
	fmt.Println(j)
	var k int
	fmt.Println(k)
	k = 22
	fmt.Println(k)

	var (
		fname string = "Scorpio"
		lname string = "King"
		empid int    = 5
	)

	fmt.Println(fname, lname, empid)

}

func ReassignVariablesExamp() {
	fmt.Println("\ninternal : ReassignVariablesExamp")
	var i int = 99
	fmt.Println(i)

	i = 24
	fmt.Println(i)

	//i:=21 // not allowed
}

var gvar1 int = 20
var gvar2 int = 30

func VariableScopeExamp() {
	fmt.Println("\ninternal : VariableScopeExamp")
	fmt.Println(gvar1, gvar2) //takes value from global declaration
	var gvar2 int = 90
	fmt.Println(gvar2) //gives more priority to variable inside function braces

}

func TypeConversion() {
	fmt.Println("\ninternal : TypeConversion")
	var i int = 20
	var j float32
	fmt.Printf("%v, %T \n", i, i) //%T is for type
	fmt.Printf("%v, %T \n", j, j)
	j = float32(i)
	fmt.Printf("%v, %T \n", i, i) //%T is for type
	fmt.Printf("%v, %T \n", j, j)

	//loss of information
	var k float32 = 75.50
	var l int
	fmt.Printf("%v, %T \n", k, k) //%T is for type
	fmt.Printf("%v, %T \n", l, l)
	l = int(k)
	fmt.Printf("%v, %T \n", k, k) //%T is for type
	fmt.Printf("%v, %T \n", l, l) //loss of information for float32 to int

	//string conversion
	var m int = 33
	var n string
	fmt.Printf("%v, %T \n", m, m)
	fmt.Printf("%v, %T \n", n, n)

	n = string(m) //it converts to unicode char
	fmt.Printf("%v, %T \n", m, m)
	fmt.Printf("%v, %T \n", n, n)

	n = strconv.Itoa(m)
	fmt.Printf("%v, %T \n", m, m)
	fmt.Printf("%v, %T \n", n, n)

}
