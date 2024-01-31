package gocontext

import (
	"context"
	"fmt"
	"time"
)

const d = 2 * time.Millisecond

func ContextWithDeadline() {
	deadline := time.Now().Add(d)

	ctx, cancel := context.WithDeadline(context.Background(), deadline)

	defer cancel()
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("Hmm")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
