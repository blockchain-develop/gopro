package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
normal test
 */
func TestSkipList1(t *testing.T) {
	sl := NewSkipList(4)
	sl.Insert(1)
	sl.Insert(3)
	sl.Insert(2)
	sl.Print()
}

/*
打印空表
 */
func TestSkipList2(t *testing.T) {
	sl := NewSkipList(4)
	sl.Print()
}

/*
失败的查询
 */
func TestSkipList3(t *testing.T) {
	sl := NewSkipList(4)
	exist := sl.Exit(1)
	assert.Equal(t, exist, false)
	sl.Print()
}

/*
失败的插入
 */
func TestSkipList4(t *testing.T) {
	sl := NewSkipList(4)
	result := sl.Insert(1)
	assert.Equal(t, result, true)
	result = sl.Insert(1)
	assert.Equal(t, result, false)
	sl.Print()
}

/*
失败的删除
 */
func TestSkipList5(t *testing.T) {
	sl := NewSkipList(4)
	result := sl.Insert(1)
	assert.Equal(t, result, true)
	result = sl.Delete(2)
	assert.Equal(t, result, false)
	sl.Print()
}

/*

 */
func TestSkipList6(t *testing.T) {
	level := 8
	dataCounter := 1000
	deleteCounter := 200
	sl := NewSkipList(level)
	data := make([]int, dataCounter)
	for i := 0;i < dataCounter;i ++ {
		data[i] = rand.Int() % (dataCounter * 10)
	}
	for i := 0;i < dataCounter;i ++ {
		if data[i] == 0 {
			data[i] ++
		}
	}

	// 插入
	for _, item := range data {
		exist := sl.Exit(item)
		result := sl.Insert(item)
		assert.Equal(t, result != exist, true)
	}

	sl.Print()

	// 元素是否都存在
	for _, item := range data {
		exist := sl.Exit(item)
		assert.Equal(t, exist, true)
	}

	// 删除前deleteCounter个元素
	for i := 0;i < deleteCounter;i ++ {
		exist := sl.Exit(data[i])
		result := sl.Delete(data[i])
		assert.Equal(t, result == exist, true)
	}

	// 元素是否都存在
	for i := 0;i < deleteCounter;i ++ {
		exist := sl.Exit(data[i])
		assert.Equal(t, exist, false)
	}

	fmt.Printf("xxxx")

	for i := deleteCounter;i < dataCounter;i ++ {
		expect := true
		for j := 0;j < deleteCounter;j ++ {
			if data[i] == data[j] {
				expect = false
			}
		}
		exist := sl.Exit(data[i])
		assert.Equal(t, exist == expect, true)
	}
	sl.Print()
}
