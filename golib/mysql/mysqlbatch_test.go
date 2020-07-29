package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestMysqlInsertBatch(t *testing.T) {
	startTime := time.Now().Unix()
	counter := 999980
	insertBatch := NewMysqlInsertBatch(DefDB, 4, "(?,?,?,?)", "INSERT INTO MysqlInsertBatch_test(name, id, url, height)")
	valueArgs := make([]interface{}, 4)
	for i := 0;i < counter;i ++ {
		valueArgs[0] = fmt.Sprintf("name:%d", i)
		valueArgs[1] = i
		valueArgs[2] = "http://www.google.com"
		valueArgs[3] = i
		err := insertBatch.Insert(valueArgs)
		if err != nil {
			panic(err)
		}
	}
	insertBatch.Close()
	endTime := time.Now().Unix()
	fmt.Printf("successful, mysqls: %d, times: %d\n", counter, endTime - startTime)
}

func TestMysqlUpdateBatch(t *testing.T) {
	counter := 999980
	{
		startTime := time.Now().Unix()
		insertBatch := NewMysqlInsertBatch(DefDB, 4, "(?,?,?,?)", "INSERT INTO MysqlUpdateBatch_test(name, id, url, height)")
		valueArgs := make([]interface{}, 4)
		for i := 0; i < counter; i ++ {
			valueArgs[0] = fmt.Sprintf("name:%d", i)
			valueArgs[1] = i
			valueArgs[2] = "http://www.google.com"
			valueArgs[3] = i
			err := insertBatch.Insert(valueArgs)
			if err != nil {
				panic(err)
			}
		}
		insertBatch.Close()
		endTime := time.Now().Unix()
		fmt.Printf("insert successful, mysqls: %d, times: %d\n", counter, endTime-startTime)
	}
	{
		startTime := time.Now().Unix()
		updateBatch := NewMysqlUpdateBatch(DefDB, 4, "(?,?,?,?)", "INSERT INTO MysqlUpdateBatch_test(name, id, url, height)", "ON DUPLICATE KEY UPDATE name=VALUES(name),url=VALUES(url)")
		valueArgs := make([]interface{}, 4)
		for i := 0; i < counter; i ++ {
			valueArgs[0] = fmt.Sprintf("name:%d", i + 10000000)
			valueArgs[1] = i
			valueArgs[2] = "http://www.baidu.com"
			valueArgs[3] = i
			err := updateBatch.Insert(valueArgs)
			if err != nil {
				panic(err)
			}
		}
		updateBatch.Close()
		endTime := time.Now().Unix()
		fmt.Printf("insert successful, mysqls: %d, times: %d\n", counter, endTime-startTime)
	}
}

func TestMysqlUpdateBatch2(t *testing.T) {
	counter := 999980
	{
		startTime := time.Now().Unix()
		insertBatch := NewMysqlInsertBatch(DefDB, 4, "(?,?,?,?)", "INSERT INTO MysqlUpdateBatch_test2(name, id, url, height)")
		valueArgs := make([]interface{}, 4)
		for i := 0; i < counter; i ++ {
			valueArgs[0] = fmt.Sprintf("name:%d", i)
			valueArgs[1] = i
			valueArgs[2] = "http://www.google.com"
			valueArgs[3] = i
			err := insertBatch.Insert(valueArgs)
			if err != nil {
				panic(err)
			}
		}
		insertBatch.Close()
		endTime := time.Now().Unix()
		fmt.Printf("insert successful, mysqls: %d, times: %d\n", counter, endTime-startTime)
	}
	{
		startTime := time.Now().Unix()
		updateBatch := NewMysqlUpdateBatch(DefDB, 4, "(?,?,?,?)", "INSERT INTO MysqlUpdateBatch_test2(name, id, url, height)", "ON DUPLICATE KEY UPDATE name=VALUES(name),url=VALUES(url)")
		valueArgs := make([]interface{}, 4)
		for i := 0; i < counter; i ++ {
			valueArgs[0] = fmt.Sprintf("name:%d", i + 10000000)
			valueArgs[1] = i
			valueArgs[2] = "http://www.baidu.com"
			valueArgs[3] = 0
			err := updateBatch.Insert(valueArgs)
			if err != nil {
				panic(err)
			}
		}
		updateBatch.Close()
		endTime := time.Now().Unix()
		fmt.Printf("update successful, mysqls: %d, times: %d\n", counter, endTime-startTime)
	}
}
