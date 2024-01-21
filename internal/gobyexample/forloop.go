package gobyexample

import "fmt"

func ForLoopExample() {

	//single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//classic initial/condition/after for loop
	for j := 4; j <= 7; j++ {
		fmt.Println(j)
	}

	//infinite loop with break
	k := 8
	for {
		fmt.Println(k)
		if k == 10 {
			break
		}
		k = k + 1
	}

	//with continue
	for l := 0; l <= 5; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}

}
