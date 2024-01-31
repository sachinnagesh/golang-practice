package gocontext

import (
	"context"
	"fmt"
	"time"
)

func ContextWithCancelExample1() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(5 * time.Second)
		cancel()

	}()
	contextScopeFunction(ctx)
}

func contextScopeFunction(ctx context.Context) {
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			printFunc()
		case <-ctx.Done():
			return
		}
	}
}

func printFunc() {
	fmt.Println("printFunc : ", time.Now())
}

func ContextWithCancelExample2() {
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(3*time.Second, cancel)
	contextScopeFunction2(ctx, 4*time.Second, "Sachin")

}

func contextScopeFunction2(ctx context.Context, d time.Duration, name string) {
	select {
	case <-time.After(d):
		fmt.Println("Name is : ", name)
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(err)

	}

}
