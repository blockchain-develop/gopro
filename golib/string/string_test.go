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
	for i := 0; i <= 200000; i++ {
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
	for i := 0; i <= 20; i++ {
		go func() {
			var which1 string
			for j := 0; j <= 10000; j++ {
				which1 += which
			}
			exitChan <- true
		}()
	}
	for i := 0; i <= 20; i++ {
		<-exitChan
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
	for i := 0; i < 200000*10; i++ {
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
	for i := 0; i < 200000*10; i++ {
		all[i] = which
	}
	strings.Join(all, "")
	timeEnd := time.Now().Unix()
	fmt.Printf("end time: %d\n", timeEnd)
}

func TestStringLower(t *testing.T) {
	aa := "0xCA20888A6F99EF8691e04D7e75f7b5faE8fa7591"
	bb := strings.ToLower(aa)
	fmt.Printf("before lower: %s, after lower: %s\n", aa, bb)
}

func TestStringLower2(t *testing.T) {
	aa := []string{"b119b3b8e5e6eeffbe754b20ee5b8a42809931fb", "b9478391eec218defa96f7b9a7938cf44e7a2fd5", "48389753b64C9e581975457332E60dC49325A653", "8F339ABc2A2a8a4D0364C7e35F892c40FBFb4BC0", "0dBbf67Fb78651D3F6407A421040f1503b486693", "89bcD91F7922126C568436841b16d036528E9714", "6514a5ebff7944099591ae3e8a5c0979c83b2571", "8c0859c191d8f100e4a3c0d8c0066c36a0c1f894", "A7d1aAc3c9bf61559c25f94132a9f801E8B5F97E", "643f3914fB8eDE03d932c79732746a8c11Ae470A", "e85631B817923487ba40263144eEfF532Cff10a2", "002E47D940dfd177dc0Fe78321E17EF84676985D"}
	for _, item := range aa {
		bb := strings.ToLower(item)
		fmt.Printf("before lower: %s, after lower: %s\n", item, bb)
	}
}

func TestStringUpper(t *testing.T) {
	aa := "8Aa2BAd3b3bC62865Ac2314FA7C8934e5C21f213"
	bb := strings.ToLower(aa)
	fmt.Printf("before lower: %s, after lower: %s\n", aa, bb)
}

func TestStringReverse(t *testing.T) {
	aa := "fd53221d8b578d599270d71a4fb2c373a2ebb75e2b0b3e0bad4b61d260956955"
	bb, _ := hex.DecodeString(aa)
	bb_len := len(bb)
	cc := make([]byte, bb_len)
	for i := 0; i < bb_len; i++ {
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

func TestSubString(t *testing.T) {
	data := "0123456789"
	len := len(data)
	if len > 5 {
		len = 5
	}
	data1 := data[0:len]
	fmt.Printf("data1: %s\n", data1)
}

func TestTemp(t *testing.T) {
	data := "2000000000"
	x, ok := new(big.Int).SetString(data, 10)
	if !ok {
		fmt.Printf("failed\n")
	} else {
		fmt.Printf("%s\n", x.String())
		fmt.Printf("%s\n", hex.EncodeToString(x.Bytes()))
	}
}

func TestTemp1(t *testing.T) {
	x := new(big.Int).SetUint64(math.MaxUint64)
	x = new(big.Int).Add(x, big.NewInt(1))
	y := new(big.Int).Mul(x, big.NewInt(1000000000))
	fmt.Printf("%s %s\n", x.String(), hex.EncodeToString(x.Bytes()))
	fmt.Printf("%s %s\n", y.String(), hex.EncodeToString(y.Bytes()))
}

func TestXXXX(t *testing.T) {
	ss := `"arbitrum","204549222"
"arbitrum","204555743"
"arbitrum","204555796"
"arbitrum","204558262"
"arbitrum","204559840"
"arbitrum","204563440"
"base","13613057"
"bitcoin","840443"
"bitcoin","840760"
"bitcoin","840761"
"bitcoin","840762"
"bitcoin","840764"
"bitcoin","840765"
"cardano","10230720"
"cardano","10230890"
"cosmos","20148338"
"doge","5185314"
"doge","5185353"
"ethereum","19729651"
"ethereum","19729670"
"ethereum","19729700"
"ethereum","19729705"
"ethereum","19729747"
"ethereum","19729778"
"ethereum","19729779"
"ethereum","19729785"
"ethereum","19729814"
"ethereum","19729834"
"ethereum","19729985"
"ethereum","19729990"
"ethereum","19730070"
"ethereum","19730089"
"injective","67329129"
"injective","67331034"
"litecoin","2673913"
"litecoin","2673914"
"litecoin","2673917"
"litecoin","2673933"
"litecoin","2673954"
"polygon","56221591"
"polygon","56222473"
"polygon","56222769"
"polygon","56223837"
"ripple","87543064"
"ripple","87543836"
"ripple","87544317"
"smartchain","38157852"
"smartchain","38158263"
"smartchain","38158771"
"smartchain","38158792"
"smartchain","38158969"
"smartchain","38159137"
"smartchain","38159543"
"smartchain","38159782"
"solana","262091622"
"solana","262092145"
"solana","262092736"
"solana","262095473"
"solana","262095487"
"solana","262095639"
"solana","262096060"
"solana","262098623"
"tron","61124150"
"tron","61124277"
"tron","61124338"
"tron","61124343"
"tron","61124367"
"tron","61124450"
"tron","61124608"
"tron","61124656"
"tron","61124817"
"tron","61125251"
"tron","61125606"
"tron","61126081"
"zcash","2482996"`
	lines := strings.Split(ss, "\n")
	curls := make([]string, 0)
	for _, line := range lines {
		items := strings.Split(line, ",")
		if len(items) != 2 {
			panic("xxx")
		}
		chain := items[0]
		blockHeight := items[1]
		//
		chain = chain[1 : len(chain)-2]
		blockHeight = blockHeight[1 : len(blockHeight)-2]
		y := `curl --location 'http://localhost:18087/api/v1/console/admin/syncer/rescan?start=%s&end=37502118&chain=ton' --header 'Content-Type: application/json' --data '{}'`
	}
}
