package goroutine

import (
	"fmt"
	"runtime/debug"
	"sort"
	"testing"
	"time"
)

/*
1. cpu任务

单线程
单线程多routine
 */
func TestPerformance_SingleThread1(t *testing.T) {
	size := 102400000
	data := make([]int, size)
	for i := 0;i < size;i ++ {
		data[i] = size - i
	}

	start := time.Now().UnixNano()
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	end := time.Now().UnixNano()
	fmt.Printf("slot time: %d\n", end - start)
}

func TestPerformance_SingleThread2(t *testing.T) {
	size := 102400000
	data := make([]int, size)
	for i := 0;i < size;i ++ {
		data[i] = size - i
	}

	start := time.Now().UnixNano()
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	end := time.Now().UnixNano()
	fmt.Printf("slot time: %d\n", end - start)
}

/*
2. cpu任务

多线程
多线程多routine
 */
func TestPerformance_MultiThread(t *testing.T) {

}

/*
3. 资源竞争的并发效率

多线程
多线程多routine
 */
func TestPerformance_ResouceRace(t *testing.T) {

}

// 测试 defer recover
func TestDferRecover(t *testing.T) {
	go Start()
	time.Sleep(100 * time.Second)
}

func Start() {
	for {
		Start1()
	}
}

func Start1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("service start, recover info: %s\n", string(debug.Stack()))
		}
	}()

	xxx := []int{1, 2}
	for i := 0;i < 3;i ++ {
		fmt.Printf("xxx: %d\n", xxx[i])
		time.Sleep(1 * time.Second)
	}
}