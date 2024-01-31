package gocontext

import (
	"context"
	"fmt"
	"time"
)

func ContextWithTimeout() {

	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	go myContextFunc(ctx)
	select {
	case <-ctx.Done():
		fmt.Println("timeout happened")
	}
	time.Sleep(3 * time.Second)

}

func myContextFunc(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		default:
			fmt.Println("Doing something")
		}
		time.Sleep(1 * time.Second)
	}

}
