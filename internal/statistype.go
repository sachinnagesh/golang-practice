package internal

import "fmt"

func TestStatictype() {
	var username string = "wagslane"
	var password string = "20947382822"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}
