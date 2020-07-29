package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"testing"
	"time"
)

func TestMultiStatements(t *testing.T) {
	startTime := time.Now().Unix()
	rounds, number := 100, 1000
	for j := 0;j < rounds;j ++ {
		var queryBuilder strings.Builder
		for i := 0; i < number; i ++ {
			queryBuilder.WriteString(fmt.Sprintf("insert into chain_info(name, id, url, height) values('%s',%d,'%s',%d);",
				fmt.Sprintf("name:%d", j * number + i), j * number + i, "xxxxxxxx", j * number + i))
		}
		_, dberr := DefDB.Exec(queryBuilder.String())
		if dberr != nil {
			fmt.Printf("failed, %s", dberr)
		}
	}
	endTime := time.Now().Unix()
	fmt.Printf("successful, mysqls: %d, times: %d", rounds * number, endTime - startTime)
}

func TestInsertBatch(t *testing.T) {
	startTime := time.Now().Unix()
	rounds, number := 100, 10000
	for j := 0;j < rounds;j ++ {
		valueStrings := make([]string, 0, number)
		valueArgs := make([]interface{}, 0, number * 4)
		for i := 0; i < number; i ++ {
			valueStrings = append(valueStrings, "(?,?,?,?)")
			valueArgs = append(valueArgs, fmt.Sprintf("name:%d", j * number + i))
			valueArgs = append(valueArgs, j * number + i)
			valueArgs = append(valueArgs, "http://www.google.com")
			valueArgs = append(valueArgs, j * number + i)
		}
		stmt := fmt.Sprintf("INSERT INTO InsertBatch_test(name, id, url, height) VALUES %s", strings.Join(valueStrings, ","))
		_, dberr := DefDB.Exec(stmt, valueArgs...)
		if dberr != nil {
			fmt.Printf("failed, %s\n", dberr)
		}
	}
	endTime := time.Now().Unix()
	fmt.Printf("successful, mysqls: %d, times: %d\n", rounds * number, endTime - startTime)
}

func TestUpdateBatch(t *testing.T) {
	rounds, number := 50, 10000
	{
		startTime := time.Now().Unix()
		for j := 0; j < rounds; j ++ {
			valueStrings := make([]string, 0, number)
			valueArgs := make([]interface{}, 0, number*4)
			for i := 0; i < number; i ++ {
				valueStrings = append(valueStrings, "(?,?,?,?)")
				valueArgs = append(valueArgs, fmt.Sprintf("name:%d", j*number+i))
				valueArgs = append(valueArgs, j*number+i)
				valueArgs = append(valueArgs, "http://www.google.com")
				valueArgs = append(valueArgs, j*number+i)
			}
			stmt := fmt.Sprintf("INSERT INTO UpdateBatch_test(name, id, url, height) VALUES %s", strings.Join(valueStrings, ","))
			_, dberr := DefDB.Exec(stmt, valueArgs...)
			if dberr != nil {
				fmt.Printf("failed, %s\n", dberr)
			}
		}
		endTime := time.Now().Unix()
		fmt.Printf("insert successful, mysqls: %d, times: %d\n", rounds*number, endTime-startTime)
	}

	// test update batch
	{
		startTime := time.Now().Unix()
		for j := 0; j < rounds; j ++ {
			valueStrings := make([]string, 0, number)
			valueArgs := make([]interface{}, 0, number*4)
			for i := 0; i < number; i ++ {
				valueStrings = append(valueStrings, "(?,?,?,?)")
				valueArgs = append(valueArgs, fmt.Sprintf("name:%d", j*number+i+100000000))
				valueArgs = append(valueArgs, j*number+i)
				valueArgs = append(valueArgs, "http://www.baidu.com")
				valueArgs = append(valueArgs, j*number+i)
			}
			stmt := fmt.Sprintf("INSERT INTO UpdateBatch_test(name, id, url, height) VALUES %s ON DUPLICATE KEY UPDATE name=VALUES(name),url=VALUES(url)", strings.Join(valueStrings, ","))
			_, dberr := DefDB.Exec(stmt, valueArgs...)
			if dberr != nil {
				fmt.Printf("failed, %s\n", dberr)
			}
		}
		endTime := time.Now().Unix()
		fmt.Printf("update successful, mysqls: %d, times: %d\n", rounds*number, endTime-startTime)
	}
}
