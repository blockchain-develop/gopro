package sleep

import (
	"fmt"
	"testing"
	"time"
)

func work1Loop() {
	fmt.Printf("start work1 loop\n")

	ticker := time.NewTicker(time.Second * 1)
	counter := 0
	for {
		select {
		case <- ticker.C:
			fmt.Printf("work1 loop counter: %d\n", counter)
			counter ++
		}
	}

	fmt.Printf("end work1 loop\n")
}

func work2Loop() {
	fmt.Printf("start work2 loop\n")

	ticker := time.NewTicker(time.Second * 1)
	counter := 0
	for {
		select {
		case <- ticker.C:
			work2Fun()
			fmt.Printf("work2 loop counter: %d\n", counter)
			counter ++
		}
	}

	fmt.Printf("end work2 loop\n")
}

func work2Fun() {

	counter := 0
	for true {
		fmt.Printf("work2 Fun counter: %d\n", counter)
		counter ++
		time.Sleep(time.Second * 1)
	}
}

var exitChain chan int

func TestSleep(t *testing.T) {
	exitChain = make(chan int)

	go work1Loop()
	go work2Loop()

	<- exitChain
	<- exitChain
}
