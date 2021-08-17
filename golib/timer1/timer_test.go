package timer1

import (
	"fmt"
	"testing"
	"time"
)

/*
任务处理时间较长，ticker被触发多次，ticker是什么行为，timer呢？

ticker: channel不会拥塞，期间ticker的事件不会出现
 */
func TestTicker(t *testing.T) {
	ti := time.NewTicker(time.Second * 1)
	triggered := false
	for {
		select {
		case <- ti.C:
			fmt.Println("ticker is trigger")
			if !triggered {
				triggered = true
				time.Sleep(time.Second * 10)
			}
		}
	}
}

func TestTimer(t *testing.T) {
	ti := time.NewTimer(time.Second * 5)
	triggered := false
	for {
		select {
		case <-ti.C:
			fmt.Println("timer is trigger")
			if !triggered {
				triggered = true
				time.Sleep(time.Second * 25)
			}
		}
	}
}

func TestTimer1(t *testing.T) {
	triggered := false
	for {
			fmt.Println("timer is trigger")
			if !triggered {
				triggered = true
				time.Sleep(time.Second * 25)
			} else {
				break
			}
		}
}

/*
如果在子逻辑中有block，那么select中的退出case还能工作吗？
 */

func TestExit(t *testing.T) {
	exit := make(chan interface{})
	queue := make(chan interface{}, 10)
	go func() {
		counter := 0
		for {
			select {
			case <-exit:
				fmt.Println("exit")
				return
			default:
			}
			counter ++
			queue <- counter
			fmt.Printf("insert %d", counter)
		}
	}()
	fmt.Printf("begin insert data to queue\n")
	time.Sleep(time.Second * 5)
	fmt.Printf("begin to exit\n")
	exit <- true
	time.Sleep(time.Second * 5)
	fmt.Printf("exit\n")
}

