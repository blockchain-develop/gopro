package _map

import (
	"fmt"
	"testing"
)

/*
测试迭代map的过程中，进行删除操作
*/
func TestMap1(t *testing.T) {
	which := make(map[int]string, 0)
	for i := 0;i < 100;i ++ {
		which[i] = fmt.Sprintf("id:%d", i)
	}
	fmt.Printf("map size: %d\n", len(which))
	for k,v := range which {
		fmt.Printf("key: %d, value: %s\n", k, v)
		delete(which, k)
	}
	fmt.Printf("map size: %d\n", len(which))
}

/*
测试迭代map的过程中，进行删除操作
*/
func TestMap2(t *testing.T) {
	which := make(map[int]string, 0)
	for i := 0;i < 100;i ++ {
		which[i] = fmt.Sprintf("id:%d", i)
	}
	fmt.Printf("map size: %d\n", len(which))
	i := 0
	for k,v := range which {
		fmt.Printf("key: %d, value: %s\n", k, v)
		delete(which, i)
		i ++
	}
	fmt.Printf("map size: %d\n", len(which))
}


/*
测试map的遍历
*/
func TestMap3(t *testing.T) {
	which := make(map[int]string, 0)
	for i := 0;i < 100;i ++ {
		which[i] = fmt.Sprintf("id:%d", i)
	}
	fmt.Printf("map size: %d\n", len(which))
	for k,v := range which {
		fmt.Printf("key: %d, value: %s\n", k, v)
	}
	fmt.Printf("map size: %d\n", len(which))
}

