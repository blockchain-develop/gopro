package context1

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("func 1 exit!\n")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("func 2 exit!\n")
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("func 3 exit!\n")
				return
			}
		}
	}()

	cancel()
	time.Sleep(time.Second * 1)
	fmt.Printf("exit!\n")
}
