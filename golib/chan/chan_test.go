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

/*
尝试想一个channel写入，如果不能写入则阻塞， 如果此时该routine还需要读入呢
A -> B 同时
B -> A
 */

func TestChan8(t *testing.T) {
	done := make(chan bool)
	r := make(chan int)
	x1 := make(chan int)
	x2 := make(chan int)
	go func() {
		for {
			select {
			case data := <- x2:
				fmt.Println("data", data)
				done <- true
				fmt.Println("finish", data)
			case data := <- r:
				fmt.Println("data", data)
				fmt.Println("finish", data)
			}
		}
	}()

	go func() {
		for {
			select {
			case data := <- x1:
				fmt.Println("data", data)
				time.Sleep(time.Second * 5)
				r <- 3
				fmt.Println("finish", data)
			case exit := <- done:
				fmt.Println("exit", exit)
			}
		}
	}()

	x1 <- 1
	x2 <- 2

	time.Sleep(time.Second * 10)
}


func TestChan9(t *testing.T) {
	done := make(chan bool)
	r := make(chan int)
	x1 := make(chan int)
	x2 := make(chan int)
	go func() {
		for {
			select {
			case data := <- x2:
				fmt.Println("data", data)
				done <- true
				fmt.Println("finish", data)
			case data := <- r:
				fmt.Println("data", data)
				fmt.Println("finish", data)
			}
		}
	}()

	writed := false
	write := func() bool {
		if writed == true {
			return true
		}
		for {
			select {
			case r <- 3:
				writed = true
				return true
			default:
				writed = false
				return false
			}
		}
	}

	go func() {
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <- ticker.C:
				write()
			case data := <- x1:
				fmt.Println("data", data)
				time.Sleep(time.Second * 5)
				write()
				fmt.Println("finish", data)
			case exit := <- done:
				fmt.Println("exit", exit)
			}
		}
	}()

	x1 <- 1
	x2 <- 2

	time.Sleep(time.Second * 10)
}
