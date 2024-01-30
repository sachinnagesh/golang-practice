package channelexamp

import (
	"fmt"
	"sync"
)

func UnbufferedChannelWithoutBlock() {
	msgc := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("In sender")
		defer wg.Done()
		msgc <- "Hi"
		fmt.Println("Sent Hi")
		msgc <- "Hello"
		fmt.Println("Sent Hello")
		msgc <- "How are you?"
		fmt.Println("Sent How are you")

	}()

	go func() {
		fmt.Println("In receiver")
		defer wg.Done()

		fmt.Printf("\nReceived : %s\n", <-msgc)
		fmt.Printf("\nReceived : %s\n", <-msgc)
		fmt.Printf("\nReceived : %s\n", <-msgc)

	}()
	wg.Wait()
}

func UnbufferedChannelDeadlockSender() {
	fmt.Println("In UnbufferedChannelDeadlockSender")
	msgc := make(chan string)
	msgc <- "hi"
	fmt.Println("THIIIIII")

}
