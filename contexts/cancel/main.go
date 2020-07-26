package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	cCtx, cFunc := context.WithCancel(ctx)

	go run(cCtx, "child")

	<-time.After(5 * time.Second)
	cFunc()
	<-time.After(5 * time.Second)
}

func run(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			println(name, "end")
			return
		case <-time.After(1 * time.Second):
			println(name, "continue")
		}
	}
}
