package time

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestTime1(t *testing.T) {
	{
		formatStr := "2006-01-02 15:04:05"
		fmt.Printf(time.Now().Format(formatStr))
	}
	{
		tt := time.Now().Unix()
		fmt.Printf("current unix time is %d or %b or %x\n", tt, tt, tt)
	}
	tt := int64(0)
	{
		tt = time.Now().UnixNano()
		fmt.Printf("current unix time is %d or %b or %x\n", tt, tt, tt)
	}
	{
		xx := time.Unix(tt/1000000000, tt%1000000000)
		formatStr := "2006-01-02 15:04:05"
		fmt.Printf(xx.Format(formatStr))
	}
}

func TestTime2(t *testing.T) {
	time_t, err := time.Parse("2006-01-02 15:04:05", "2020-09-25 22:11:33")
	if err != nil {
		panic(err)
	}
	time_unix := time_t.Unix()
	time_str := time_t.Format("2006-01-02 15:04:05")
	fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
}

func TestTime3(t *testing.T) {
	time_t := time.Unix(1598583600, 0)
	time_unix := time_t.Unix()
	time_str := time_t.Format("2006-01-02 15:04:05")
	fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
}

func TestTimeTransfer(t *testing.T) {
	{
		time_t := time.Now()
		time_unix := time_t.Unix()
		time_str := time_t.Format("2006-01-02 15:04:05")
		fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
	}

	{
		time_t := time.Unix(1593332467, 0)
		time_unix := time_t.Unix()
		time_str := time_t.Format("2006-01-02 15:04:05")
		fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
	}

	{
		time_t, err := time.Parse("2006-01-02 15:04:05", "2020-06-28 16:21:07")
		if err != nil {
			panic(err)
		}
		time_unix := time_t.Unix()
		time_str := time_t.Format("2006-01-02 15:04:05")
		fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
	}
}

func TestTimeDays(t *testing.T) {
	time_t := time.Unix(1593388800, 0)
	time_unix := time_t.Unix()
	time_str := time_t.Format("2006-01-02")

	time_t_new, err := time.Parse("2006-01-02", time_str)
	if err != nil {
		panic(err)
	}
	time_unix_new := time_t_new.Unix()
	time_str_new := time_t_new.Format("2006-01-02")
	fmt.Printf("old time, unix: %d, str: %s, new time, unix: %d, str: %s\n", time_unix, time_str, time_unix_new, time_str_new)
}

func TestTimeElement(t *testing.T) {
	time_t := time.Now()
	fmt.Printf("year: %d, month: %d, day: %d, hour: %d, minute: %d, second: %d\n",
		time_t.Year(), time_t.Month(), time_t.Day(), time_t.Hour(), time_t.Minute(), time_t.Second())
}

func  DayOfTimeDown(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	return uint32(end_t_new.Unix())
}

func  DayOfTimeUp(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.Unix())
	if t > time_t_unix {
		time_t_unix = uint32(end_t_new.AddDate(0, 0, 1).Unix())
	}
	return time_t_unix
}

func  DayOfTimeAddOne(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.AddDate(0, 0, 1).Unix())
	return time_t_unix
}

func  DayOfTimeSubOne(t uint32) uint32 {
	end_t := time.Unix(int64(t), 0)
	end_t_new, _ := time.Parse("2006-01-02", end_t.Format("2006-01-02"))
	time_t_unix := uint32(end_t_new.AddDate(0, 0, -1).Unix())
	return time_t_unix
}

func TestTimeSecond(t *testing.T) {
	tt := time.Now()
	tt_str := tt.Format("2006-01-02 15:04:05")
	fmt.Printf("now: %s \n", tt_str)
	new_tt := time.Unix((tt.Unix() / 60) * 60, 0)
	new_tt_str := new_tt.Format("2006-01-02 15:04:05")
	fmt.Printf("now: %s \n", new_tt_str)
}

func TestUintMax(t *testing.T) {
	a := math.MaxUint32
	fmt.Printf("max uint32: %d\n", a)
}

func TestBigInt(t *testing.T) {
	aa := big.NewInt(100)
	bb := big.NewInt(10)
	//aa.Mod(aa, bb)
	aa.Div(aa, bb)
	fmt.Printf("aa: %d, bb: %d\n", aa.Int64(), bb.Int64())
}

func TestMillisecond(t *testing.T) {
	t1 := uint64(time.Second.Milliseconds()) * 5
	t2 := t1 * uint64(time.Millisecond)
	fmt.Println("start")
	time.Sleep(time.Duration(t2))
	fmt.Println("exit")
}

func TestNanosecond(t *testing.T) {
	t1 := time.Now().UnixNano()
	t2 := time.Now().UnixNano()
	t3 := time.Now().UnixNano()
	fmt.Printf(" t1: %d\n t2: %d\n t3: %d\n", t1, t2, t3)
	t4 := t3 >> 32
	t5 := time.Now().Unix()
	fmt.Printf(" t4: %d\n, t5: %d\n", t4, t5)
}

func TestXXXX(t *testing.T) {
	index := uint(2)
	id  := int64(0)
	{
		tt := time.Now().Unix()
		id = int64(index) << 32 | tt
		fmt.Printf("id: %d\n", id)
	}
	{
		index = uint(id >> 32)
		fmt.Printf("index: %d\n", index)
	}
}

func Test11(t *testing.T) {
	version := "/Dash Core:18.0.1/"
	version = strings.TrimPrefix(
		strings.TrimSuffix(version, "/"),
		"/Dash Core:",
	)
	switch {
	case version < "0.19.0":
		fmt.Printf("xx")
	default:
		fmt.Printf("bb")
	}
}
