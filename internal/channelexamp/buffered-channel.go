package channelexamp

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*
**
DeadLock

A deadlock in a buffered channel in Go occurs when a goroutine is trying to send data to a channel that is already full,
or when a goroutine is trying to receive data from an empty channel.

When a goroutine tries to send data to a full channel, it blocks until there is space available in the channel buffer. If another goroutine is trying to receive data from the same channel
and is also blocked, a deadlock occurs.

Similarly, when a goroutine tries to receive data from an empty channel,it blocks until there is data available in the channel buffer. If another goroutine is trying to send data to the same channel
and is also blocked, a deadlock occurs.

**
*/
func BufferedChannelWithoutBlock() {

	msgc := make(chan int, 2) //creating a buffered channel with capacity 2

	msgc <- 3           //puting first element
	msgc <- 5           //putting second element
	fmt.Println(<-msgc) //get first element
	fmt.Println(<-msgc) //get second element

}

// A deadlock in a buffered channel in Go occurs when a goroutine is trying to send data to a channel that is already full,
func BufferedChannelDeadLockSender() {
	msgc := make(chan string, 2)

	msgc <- "hi"
	msgc <- "Hello"
	fmt.Println("BufferedChannelDeadLock : line after this creates deadlock !!!")
	msgc <- "How are you?" //this creates deadlock as there is no space to put value as it's not recevied
	fmt.Println(<-msgc)
	fmt.Println(<-msgc)
	fmt.Println(<-msgc)
}

// when a goroutine is trying to receive data from an empty channel
func BufferedChannelDeadLockReceiver() {
	msgc := make(chan string, 3)

	msgc <- "Hi"
	msgc <- "Hello"
	fmt.Println(<-msgc)
	fmt.Println(<-msgc) //after this channel is empty
	fmt.Println("line after this creates deadlocak !!!")
	fmt.Println(<-msgc)
	msgc <- "how are you?"
}

func ChannelWithGoRoutineNoBlocking() {

	msgc := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2) //no of goroutines

	go func() {
		msg1 := "Hi"
		msg2 := "Hello"
		defer wg.Done()
		fmt.Printf("\nSending message : %s and %s", msg1, msg2)
		msgc <- msg1
		msgc <- msg2
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		rmsg1 := <-msgc
		rmsg2 := <-msgc
		fmt.Printf("\nReceived message : %s and %s", rmsg1, rmsg2)

	}()

	wg.Wait()

}

func ChannelWithGoRoutineWithBlock() {
	msgc := make(chan string, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("In sender")
		msgc <- "Hi"
		msgc <- "Hello"
		fmt.Println("In sender : channel is full so this is blocked now")
		msgc <- "How are you?"
		fmt.Println("In sender : sent How are you? to channel. Receiver must have received message !!")

		fmt.Println("In sender : sending message : I am Sachin")
		for i := 0; i < 6; i++ {
			fmt.Printf("\nSending iteration : %d", i)
			msgc <- "I am Sachin " + strconv.Itoa(i)
		}

		defer wg.Done()
	}()

	go func() {
		fmt.Println("In receiver")
		time.Sleep(time.Second * 4)
		rmsg1 := <-msgc
		rmsg2 := <-msgc
		fmt.Printf("\nReceived msg : %s and %s", rmsg1, rmsg2)
		for i := 0; i < 6; i++ {
			fmt.Printf("\nReceived message : %s", <-msgc)
			fmt.Printf("\nReceiving Iteration : %d", i)
		}

		defer wg.Done()
	}()
	fmt.Println("Before wg.Wait")
	wg.Wait()
	fmt.Println("This should be last sentence !!!")
}
