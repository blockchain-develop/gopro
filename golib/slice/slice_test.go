package slice

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func dumpSlice1(name string, s []byte) {
	len := len(s)
	cap := cap(s)
	p := unsafe.Pointer(&s)
	h := *(*reflect.SliceHeader)(p)
	fmt.Printf("%s, slice: %v, %p, len: %d, cap: %d, data: %p\n", name, s, &s, len, cap, unsafe.Pointer(h.Data))
}

func dumpSlice(name string, s *[]byte) {
	len := len(*s)
	cap := cap(*s)
	p := unsafe.Pointer(s)
	h := *(*reflect.SliceHeader)(p)
	fmt.Printf("%s, slice: %v, %p, len: %d, cap: %d, data: %p\n", name, *s, s, len, cap, unsafe.Pointer(h.Data))
}

/*
调用函数，拷贝slice对象，但并不拷贝slice中的data，即浅拷贝
slice对象位于栈上，在栈上拷贝对象
slice中的data位于堆上，并没有拷贝
*/
func TestSlice1(t *testing.T) {
	which := []byte("1234567")
	fmt.Printf("source, slice: %p\n", &which)
	dumpSlice("dump reference", &which)
	dumpSlice1("dump copy    ", which)
}

/*
slice的data是可以被多个slice共享的，但并没有COW，data被修改，所有引用该data的slice数据都会更新
 */
func TestSlice2(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("1234567")...)
	which1 := which
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which1[0] = '2'
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
}

/*
slice的data被多个slice共享，多个slice交易追加数据会出现数据覆写问题
 */
func TestSlice3(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("123")...)
	which1 := which
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which = append(which, []byte("45")...)
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which1 = append(which1, []byte("67")...)
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
}

/*
slice的data被多个slice共享，多个slice交易追加数据会出现数据覆写问题，覆写的另一个例子
 */
func TestSlice4(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("123456")...)
	which1 := which[1:3]
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which1 = append(which1, []byte("123")...)
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
}

/*
slice的data被多个slice共享，但在某个slice发生扩容后，其data区域不再和原来的slice一样，修改数据对原有的slice不再有影响
 */
func TestSlice5(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("123")...)
	which1 := which
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which = append(which, []byte("45")...)
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	which1 = append(which1, []byte("123456")...)
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
}

/*
slice的data共享但没有cow机制的综合例子
 */
func TestSlice6(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("abcdef")...)
	which1 := which
	which2 := which[1:3]
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	dumpSlice("dump which2     ", &which2)
	which2 = append(which2, 'x')
	which1[2] = 'x'
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	dumpSlice("dump which2     ", &which2)
	which1[2] = 'y'
	which = append(which, []byte("ghij")...)
	which2 = append(which2, 'z')
	which1[2] = 'z'
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1     ", &which1)
	dumpSlice("dump which2     ", &which2)
}

/*
扩容
预估新数据需要的总容量TP
当前的容量CP
如果TP > CP, 则扩容后新cap = 256向上取整(TP)
如果CP < 1024, 则扩容后的新cap = (2 * CP)
如果CP > 1024, 则扩容后的新cap = 256向上取整(1.25 * CP)
*/
func TestSlice7(t *testing.T) {
	which := make([]byte, 0)
	for i := 0;i < 500;i ++ {
		which = append(which, []byte("abcdefgh")...)
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
	}
}

func TestSlice8(t *testing.T) {
	which := make([]byte, 0)
	for i := 0;i < 500;i ++ {
		which = append(which, []byte("abcde")...)
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
	}
}

func TestSlice9(t *testing.T) {
	which := make([]byte, 0)
	for i := 0;i < 128;i ++ {
		which = append(which, []byte("abcdefgh")...)
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
	}
	which = append(which, which...)
	fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
}

func TestSlice10(t *testing.T) {
	which := make([]byte, 0)
	for i := 0;i < 8;i ++ {
		which = append(which, []byte("abcdefgh")...)
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
		dumpSlice("dump which      ", &which)
		which = which[7:]
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
		dumpSlice("dump which      ", &which)
		fmt.Printf("=====================================================\n")
	}
}

func TestSlice11(t *testing.T) {
	which := make([]byte, 0)
	for i := 0;i < 100;i ++ {
		which = append(which, []byte("abcdefgh")...)
		fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
	}
	which = which[512:]
	fmt.Printf("size: %d\t, cap: %d\t\n", len(which), cap(which))
}

/*
指定容量，
从一个slice append到当前slice，进行data的拷贝操作
重置slice，使用[0:0]进行slice重置，
*/
func TestSlice12(t *testing.T) {
	which1 := make([]byte, 0, 1024)
	which2 := make([]byte, 0)
	which2 = append(which2, []byte("abcdefgh")...)
	dumpSlice("dump which1      ", &which1)
	dumpSlice("dump which2      ", &which2)

	which1 = append(which1, which2...)
	dumpSlice("dump which1      ", &which1)
	dumpSlice("dump which2      ", &which2)

	which2[0] = which2[7]
	which1 = append(which1, which2...)
	dumpSlice("dump which1      ", &which1)
	dumpSlice("dump which2      ", &which2)

	which1 = which1[0:0]
	dumpSlice("dump which1      ", &which1)

	which1 = append(which1, which2...)
	dumpSlice("dump which1      ", &which1)
}

/*
slice的赋值和append的区别
赋值操作导致两个slice的data是共享的
append操作的两个slice的data不是共享的
*/
func TestSlice13(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("123456")...)
	which1 := make([]byte, 0)
	which1 = append(which1, which...)
	which2 := which
	dumpSlice("dump which      ", &which)
	dumpSlice("dump which1      ", &which1)
	dumpSlice("dump which2      ", &which2)
}

/*
测试指定容量的情况下，元素溢出情况
元素溢出，继续扩容
*/
func TestSlice14(t *testing.T) {
	which1 := make([]byte, 0, 16)
	which1 = append(which1, []byte("abcdefgh")...)
	dumpSlice("dump which1      ", &which1)

	which1 = append(which1, []byte("abcdefgh")...)
	dumpSlice("dump which1      ", &which1)

	which1 = append(which1, []byte("abcdefgh")...)
	dumpSlice("dump which1      ", &which1)
}

/*
测试指定容量的情况下，缩容情况
*/
func TestSlice15(t *testing.T) {
	which1 := make([]byte, 0, 32)
	for i := 0;i < 128;i ++ {
		which1 = append(which1, []byte("abcdefgh")...)
	}
	dumpSlice("dump which1      ", &which1)

	which1 = which1[0:0]
	dumpSlice("dump which1      ", &which1)
}

/*
测试指定容量的情况下，缩容情况
如果slice的data有很多多余的无法再被使用的空间，则重新构造data区
*/
func TestSlice16(t *testing.T) {
	which1 := make([]byte, 0, 32)
	for i := 0;i < 128;i ++ {
		which1 = append(which1, []byte("abcdefgh")...)
	}
	dumpSlice("dump which1      ", &which1)

	which1 = which1[1000:1024]
	dumpSlice("dump which1      ", &which1)
}

/*
测试slice的range，在range最后退出的时候，其index值
*/
func TestSlice17(t *testing.T) {
	which := make([]byte, 0)
	which = append(which, []byte("abcd")...)

	var index int
	var char byte
	for index, char = range which {

	}
	fmt.Printf("index: %d, char: %d\n", index, char)
}

/*
测试在指定容量和指定大小两种场景下，设置slice的元素
*/
func TestSlice18(t *testing.T) {
	which := make([]byte, 10, 10)
	for i := 0;i < 5;i ++ {
		which[i] = byte(i)
	}
	for i := 0;i < 5;i ++ {
		which[10 + i] = byte(i)
	}
}

func TestSlice19(t *testing.T) {
	which := make([]byte, 10, 20)
	for i := 0;i < 5;i ++ {
		which[i] = byte(i)
	}
	for i := 0;i < 5;i ++ {
		which[10 + i] = byte(i)
	}
}

/*
以下测试指针的slice和对象slice
*/
type AAA struct {
	aaa    int
	bbb    int
}

func TestSlice20(t *testing.T) {
	test1 := []*AAA{
		&AAA {
			aaa: 0,
			bbb: 0,
		},
		&AAA {
			aaa: 1,
			bbb: 1,
		},
	}
	test2 := make([]*AAA, 0, 2)
	test2 = append(test2, test1...)
	for _, item := range test2 {
		item.aaa = item.aaa + 10
		item.bbb = item.bbb + 10
	}

	fmt.Printf("test1:\n")
	for _, item := range test1 {
		fmt.Printf("aaa: %d, bbb: %d\n", item.aaa, item.bbb)
	}
	fmt.Printf("test2:\n")
	for _, item := range test2 {
		fmt.Printf("aaa: %d, bbb: %d\n", item.aaa, item.bbb)
	}
}

func TestSlice21(t *testing.T) {
	a := [2]int{5, 6}
	b := [2]int{5, 6}

	// ①
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

	/*
	// ②
	if a[:] == b[:] {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
	*/
}

func TestSlice22(t *testing.T) {
	var xxxx [32]byte
	copy(xxxx[0:2], []byte("xx"))
	fmt.Printf("xxxx: %v\n", xxxx)
}

func TestSlice23(t *testing.T) {
	var aaa [32]byte
	var bbb [32]byte
	aaa[0] = 1
	fmt.Printf("a == b ? %v", aaa == bbb)
	/*
	ccc := make([]byte, 0)
	ddd := make([]byte, 0)
	fmt.Printf("c == d ? %v", ccc == ddd)
	 */
}
