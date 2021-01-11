package string

import (
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
	"unsafe"
)

func pointer(s *string) uintptr {
	p := unsafe.Pointer(s)
	h := *(*reflect.StringHeader)(p)
	return h.Data
}

func dumpString1(s string) {
	len := len(s)
	data := pointer(&s)
	fmt.Printf("dump  , string: %s, %p, len: %d, data: %p\n", s, &s, len, unsafe.Pointer(data))
}

func dumpString(s *string) {
	len := len(*s)
	data := pointer(s)
	fmt.Printf("dump  , string: %s, %p, len: %d, data: %p\n", *s, s, len, unsafe.Pointer(data))
}

func dumpSlice(s *[]byte) {
	len := len(*s)
	cap := cap(*s)
	p := unsafe.Pointer(s)
	h := *(*reflect.SliceHeader)(p)
	fmt.Printf("dump  , slice: %v, %p, len: %d, cap: %d, data: %p\n", *s, s, len, cap, unsafe.Pointer(h.Data))
}

/*
调用函数，拷贝string对象，但并不拷贝string中的data，即浅拷贝
string对象位于栈上，在栈上拷贝对象
string中的data位于堆上，并没有拷贝
*/
func TestString1(t *testing.T) {
	which := "hello,world!"
	fmt.Printf("source, string: %s, %p\n", which, &which)
	dumpString1(which)
	dumpString(&which)
}

/*
通过切片方式获取一个新string，其data指向的数据区和原string一样，但指针有偏移
通过位置索引方式取得字符并生成一个新的strig后，其data指向的数据区和原string不一样
 */
func TestString2(t *testing.T) {
	which := "hello,world!"
	dumpString(&which)
	which2 := which[6:8]
	dumpString(&which2)
	which1 := string(which[6])
	dumpString(&which1)
}

/*
string被修改，只是修改其指向的data，go string具有不可变性
 */
func TestString3(t *testing.T) {
	which := "hello,world!"
	which1 := which
	dumpString(&which)
	dumpString(&which1)
	which = "how are you!"
	dumpString(&which)
	dumpString(&which1)
}

/*
string转换的切片的data并不是直接指向的string的data区域，而是将string的data进行了拷贝
*/
func TestString4(t *testing.T) {
	which := "hello,world!"
	dumpString(&which)
	which1 := []byte(which)
	dumpSlice(&which1)
}

/*
string比较操作，string的data指向不一样，单内容一样
 */
func TestString5(t *testing.T) {
	which := "hello,world!"
	which1 := "hello,world!123456"
	which1 = which1[0:12]
	dumpString(&which)
	dumpString(&which1)
	if which == which1 {
		fmt.Printf("same!\n")
	} else {
		fmt.Printf("different!\n")
	}
}

/*
由于string的不可变性，导致如下的操作性能非常的低下，每次的+=操作都导致在堆上分配内存和数据拷贝
以下是单routine和多routine的版本
 */
func TestString6(t *testing.T) {
	which := "hello,world!"
	timeBegin := time.Now().Unix()
	fmt.Printf("begin time: %d\n", timeBegin)
	var which1 string
	for i := 0;i <= 200000;i ++ {
		which1 += which
	}
	timeEnd := time.Now().Unix()
	fmt.Printf("end time: %d\n", timeEnd)
}

func TestString7(t *testing.T) {
	which := "hello,world!"
	timeBegin := time.Now().Unix()
	fmt.Printf("begin time: %d\n", timeBegin)
	exitChan := make(chan bool, 20)
	for i := 0;i <= 20;i ++ {
		go func() {
			var which1 string
			for j := 0;j <= 10000;j ++ {
				which1 += which
			}
			exitChan <- true
		}()
	}
	for i := 0;i <= 20;i ++ {
		<- exitChan
	}
	timeEnd := time.Now().Unix()
	fmt.Printf("end time: %d\n", timeEnd)
}

/*
自动扩容的builder
 */
func TestString8(t *testing.T) {
	which := "hello,world!"
	timeBegin := time.Now().Unix()
	fmt.Printf("begin time: %d\n", timeBegin)
	var builder strings.Builder
	for i := 0;i < 200000*10;i ++ {
		builder.WriteString(which)
	}
	timeEnd := time.Now().Unix()
	fmt.Printf("end time: %d\n", timeEnd)
}

/*
join首先会计算处需要的总长度，一次分配内存，然后拷贝数据
 */
func TestString9(t *testing.T) {
	which := "hello,world!"
	timeBegin := time.Now().Unix()
	fmt.Printf("begin time: %d\n", timeBegin)
	all := make([]string, 200000*10)
	for i := 0;i < 200000*10;i ++ {
		all[i] = which
	}
	strings.Join(all, "")
	timeEnd := time.Now().Unix()
	fmt.Printf("end time: %d\n", timeEnd)
}

func TestStringLower(t *testing.T) {
	aa := "1269d9940a2bfc5aC13c759E7ef1E35FEc7278f6"
	bb := strings.ToLower(aa)
	fmt.Printf("before lower: %s, after lower: %s\n", aa, bb)
}

func TestStringUpper(t *testing.T) {
	aa := "D8aE73e06552E270340b63A8bcAbf9277a1aac99"
	bb := strings.ToLower(aa)
	fmt.Printf("before lower: %s, after lower: %s\n", aa, bb)
}

func TestStringReverse(t *testing.T) {
	aa := "b6cb731f90cefebbd4f9cedd0cf56bd1e21967f4"
	bb, _ := hex.DecodeString(aa)
	bb_len := len(bb)
	cc := make([]byte, bb_len)
	for i := 0;i < bb_len;i ++ {
		cc[bb_len-1-i] = bb[i]
	}
	fmt.Printf("old string: %s\n", aa)
	fmt.Printf("new string: %s\n", hex.EncodeToString(cc))
}

func HexReverse(arr []byte) []byte {
	l := len(arr)
	x := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		x = append(x, arr[i])
	}
	return x
}

func HexStringReverse(value string) string {
	aa, _ := hex.DecodeString(value)
	bb := HexReverse(aa)
	return hex.EncodeToString(bb)
}

func TestStringDecode(t *testing.T) {
	xx := "005007"
	//bb, _ := hex.DecodeString(xx)
	cc, _ := strconv.ParseUint(xx, 16, 64)
	dd, _ := strconv.ParseUint(HexStringReverse(xx), 16, 64)
	fmt.Printf("xxxx: %d, %d\n", cc, dd)
}

func TestDataIncode(t *testing.T) {
	xxx := big.NewInt(999)
	xxx.Mul(xxx, big.NewInt(1000000000000000000))
	fmt.Printf("xxx: %s\n", hex.EncodeToString(xxx.Bytes()))
	fmt.Printf("amount: %d\n", xxx.Uint64())
	fmt.Printf("max uint64: %d, int64: %d", uint64(math.MaxUint64), int64(math.MaxInt64))
}

func TestStringDecode2(t *testing.T) {
	xx := "005007"
	bb, err := hex.DecodeString(xx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("decode: %s\n", bb)
}

func TestStringJoin(t *testing.T) {
	test := []string{"aaaa", "bbbb", "cccc"}
	res := strings.Join(test, ",")
	fmt.Printf("result: %s\n", res)
}