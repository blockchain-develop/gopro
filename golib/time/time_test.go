package time

import (
	"fmt"
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
}

func TestTime2(t *testing.T) {
	time_t, err := time.Parse("2006-01-02 15:04:05", "2020-06-30 16:21:07")
	if err != nil {
		panic(err)
	}
	time_unix := time_t.Unix()
	time_str := time_t.Format("2006-01-02 15:04:05")
	fmt.Printf("current time, unix: %d, str: %s\n", time_unix, time_str)
}

func TestTime3(t *testing.T) {
	time_t := time.Unix(1592956800, 0)
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

func outputCrossChainTxStatus(status []*CrossChainTxStatus, start uint32, end uint32, total uint32) []*CrossChainTxStatus {
	current_txnumber := total
	current_tt := uint32(0)
	status_new := make([]*CrossChainTxStatus, 0)
	for _, s := range status {
		status_new = append(status_new, &CrossChainTxStatus{
			TT: s.TT,
			TxNumber: current_txnumber,
		})
		current_txnumber = current_txnumber - s.TxNumber
		current_tt = s.TT
	}
	status_new = append(status_new, &CrossChainTxStatus{
		TT: DayOfTimeSubOne(current_tt),
		TxNumber: current_txnumber,
	})

	status_new1 := make([]*CrossChainTxStatus, 0)
	if len(status_new) != 0 {
		current_txnumber = status_new[0].TxNumber
		current_tt := status_new[0].TT
		for _, s := range status_new {
			for s.TT < current_tt {
				status_new1 = append(status_new1, &CrossChainTxStatus{
					TT:       current_tt,
					TxNumber: current_txnumber,
				})
				current_tt = DayOfTimeSubOne(current_tt)
			}

			current_txnumber = s.TxNumber
			status_new1 = append(status_new1, &CrossChainTxStatus{
				TT:       current_tt,
				TxNumber: current_txnumber,
			})
			current_tt = DayOfTimeSubOne(current_tt)
		}
	}

	status_new = make([]*CrossChainTxStatus, 0)
	if len(status_new1) != 0 {
		ss := status_new1[len(status_new1) - 1]
		current_txnumber = ss.TxNumber
		current_tt := DayOfTimeDown(start)
		for current_tt < ss.TT {
			status_new = append(status_new, &CrossChainTxStatus{
				TT:       current_tt,
				TxNumber: current_txnumber,
			})
			current_tt = DayOfTimeAddOne(current_tt)
		}
		for i := 0;i < len(status_new1);i ++ {
			bb := status_new1[len(status_new1) - 1 - i]
			status_new = append(status_new, bb)
			current_tt = bb.TT
			current_txnumber = bb.TxNumber
		}
		for current_tt < DayOfTimeUp(end) {
			status_new = append(status_new, &CrossChainTxStatus{
				TT:       current_tt,
				TxNumber: current_txnumber,
			})
			current_tt = DayOfTimeAddOne(current_tt)
		}
	}
	return status_new
}

type CrossChainTxStatus struct {
	TT        uint32    `json:"timestamp"`
	TxNumber  uint32    `json:"txnumber"`
}

func TestOutputCrossChainTxStatus(t *testing.T) {
	status := make([]*CrossChainTxStatus, 0)
	status = append(status, &CrossChainTxStatus{
		TT : 1593014400,
		TxNumber: 7,
	})
	outputCrossChainTxStatus(status, 1592274867, 1593534067, 7)
}
