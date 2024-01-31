package gocontext

import (
	"context"
	"fmt"
	"time"
)

/**
Making HTTP requests is one of the withValue function's most popular use cases.
With the Context() method, the net/Http package in Go provides internal logic
that builds a context for each request.
**/

func CreateContextWithValue() {

	ctx := context.Background()

	ctx = context.WithValue(ctx, "key1", "value1")
	printconextkeyandvalue(ctx)
	time.NewTicker(500 * time.Millisecond)

}

func printconextkeyandvalue(ctx context.Context) {
	fmt.Println("Context value with key : key1 is ", ctx.Value("key1"))

}
