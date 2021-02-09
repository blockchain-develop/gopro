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

func TestChan3(t *testing.T) {
	exit := make(chan bool, 0)
	go func() {
		for {
			//time.Sleep(5 * time.Second)
			select {
			case <- exit:
				fmt.Printf("aaaa\n")
				//time.Sleep(5 * time.Second)
				//return
			}
		}
	}()
	//exit <- true
	close(exit)
	//do something after routine
	fmt.Printf("bbbb\n")
	time.Sleep(time.Second * 5)
}

func TestChan4(t *testing.T) {
	exit := make(chan bool, 0)
	go func() {
		for {
			//time.Sleep(5 * time.Second)
			select {
			case v, _ := <- exit:
				fmt.Printf("aaaa: %v\n", v)
				//time.Sleep(5 * time.Second)
				//return
			}
		}
	}()
	exit <- true
	exit <- true
	exit <- true
	//
	close(exit)


	//do something after routine
	fmt.Printf("bbbb\n")
	time.Sleep(time.Second * 5)
}
