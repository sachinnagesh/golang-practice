package gocontext

import (
	"context"
	"fmt"
)

/*
*

	The two ways to create a Context:
	- context.Background()
	- context.TODO()
	Both functions return a non-nil, empty context. The only time TODO is used instead of Background
	is when the implementation is unclear or the context is not yet known.

*
*/
func CreateContextWithTODO() {
	ctx := context.TODO()
	myfunc(ctx)
}

func CreateContextWithBackground() {
	ctx := context.Background()
	myfunc(ctx)
}

func myfunc(ctx context.Context) {
	fmt.Println("function with context param")

}
