package _chan

import (
	"fmt"
	"testing"
	"time"
)

func work(exit *chan int) {
	for i:= 0;i < 10;i ++ {
		time.Sleep(time.Second * 1)
	}
	<- *exit
}

func TestChan(t *testing.T) {
	exit := make(chan int)
	go work(&exit)
	exit <- 1
	fmt.Printf("Test Chan")
}

func TestChan2(t *testing.T) {
	exit := make(chan int, 1)
	exit <- 1

	work(&exit)

	select {
	case exit <- 1:
		fmt.Printf("xxxx")
	}

}
