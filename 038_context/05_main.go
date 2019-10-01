package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Make sure all paths cancel the context to avoid context leak

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
	fmt.Println(time.Now())
	time.Sleep(1 * time.Minute)

}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}

	}()

	return ch
}

//https://rakyll.org/leakingctx/
