package schedule

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/*
启动两个roroutine，可以正确调度
这是一个正常的case，这两个routine可能会被GMP调度在不同的线程上运行。
 */
func TestSchedule1(t *testing.T) {
	go func() {
		for true {
			fmt.Printf("BBB\n")
			time.Sleep(1 * time.Second)
		}
	}()
	for true {
		fmt.Printf("AAA\n")
		time.Sleep(1 * time.Second)
	}
}

/*
限制启动一个线程，那么这两个routine只能被GMP调度在同一个线程上运行。
sleep和printf是函数调用，可以做伪抢占式调度，上面的二个方案可以实现，这两个routine都可以被调度执行
 */
func TestSchedule2(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		for true {
			fmt.Printf("BBB\n")
			time.Sleep(1 * time.Second)
		}
	}()
	for true {
		fmt.Printf("AAA\n")
		time.Sleep(1 * time.Second)
	}
}

/*
限制启动一个线程，那么这两个routine只能被GMP调度在同一个线程上运行。
没有函数调用情况下，用户态的协程只能在抢占式调度下被切换，上面的第三个方式可以实现，这两个routine照样可以被调度执行。
只有go 1.14有信号实现的抢占式调度，所以以下代码在go 1.14在，两个routine可以被调度，会panic，而go 1.13和之前版本没有信号实现的抢占式调度，不会panic。
 */
func TestSchedule3(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		panic("Can not here?")
	}()
	for true {
		continue
	}
}

/*
启动二个线程，那么这两个routine可能会被GMP调度在二个线程上运行。
没有函数调用情况下，没有抢占式调度的情况下，这两个routine可以执行，因为他们可以被分配在两个线程上运行，线程调度是操作系统完成的，有抢占式调度机制的。
以下代码在go 1.13下也会panic
 */
func TestSchedule4(t *testing.T) {
	runtime.GOMAXPROCS(1)
	go func() {
		panic("Can not here?")
	}()
	for true {
		continue
	}
}

