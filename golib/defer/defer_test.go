package _defer

import (
	"fmt"
	"testing"
)

/*
1. 作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值
2. 多个defer的执行顺序为“后进先出”
3. defer、return、返回值三者的执行逻辑应该是：return最先执行，return负责将结果写入返回值中；接着defer开始执行一些收尾工作；最后函数携带当前返回值退出。
4. 如果函数的返回值是无名的（不带命名返回值），则go语言会在执行return的时候会执行一个类似创建一个临时变量作为保存return值的动作，而有名返回值的函数，由于返回值在函数定义的时候已经将该变量进行定义，在执行return的时候会先执行返回值保存操作，而后续的defer函数会改变这个返回值(虽然defer是在return之后执行的，但是由于使用的函数定义的变量，所以执行defer操作后对该变量的修改会影响到return的值
*/
func test1() int {
	var i int
	defer func() {
		i ++
		fmt.Printf("aaa: %d\n", i)
	}()
	defer func() {
		i ++
		fmt.Printf("bbb: %d\n", i)
	}()
	return i
}

func TestDefer1(t *testing.T) {
	i := test1()
	fmt.Printf("ccc: %d\n", i)
}

func test2() (i int) {
	defer func() {
		i ++
		fmt.Printf("aaa: %d\n", i)
	}()
	defer func() {
		i ++
		fmt.Printf("bbb: %d\n", i)
	}()
	return i
}

func TestDefer2(t *testing.T) {
	i := test2()
	fmt.Printf("ccc: %d\n", i)
}
