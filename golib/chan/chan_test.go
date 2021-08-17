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

func TestChan5(t *testing.T) {
	exit := make(chan bool)
	go func() {
		for {
			select {
			case <- exit:
				fmt.Printf("aaa\n")
			}
		}
	}()

	exit <- true
	exit <- true
	exit <- true

	fmt.Printf("bbb\n")
	time.Sleep(time.Second * 5)
}

/*
尝试向一个channel写入，如果不能写则立马返回，不要阻塞
 */
func TestChan6(t *testing.T) {
	data := make(chan interface{}, 5)
	for {
		select {
			case data <- true:
				fmt.Println("write")
		}
	}
}

func TestChan7(t *testing.T) {
	data := make(chan interface{}, 5)
	for {
		select {
		case data <- true:
			fmt.Println("write")
		default:
			fmt.Printf("continue")
		}
	}
}
