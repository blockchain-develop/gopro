package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

type Contract struct {
	Address         string  `json:"address"`
	Chain           uint32  `json:"chain"`
}
type Token struct {
	Name            string  `json:"name"`
	Contracts       []Contract  `json:"contracts"`
}
type AAAA struct {
	Tokens          []Token     `json:"tokens"`
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", fileName, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Printf("File %s close error %s", fileName, err)
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", fileName, err)
	}
	return data, nil
}

func TestJson(t *testing.T) {
	fileName := "./test1.json"
	data, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}

	var aaa AAAA
	err = json.Unmarshal(data, &aaa)
	if err != nil {
		fmt.Printf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}
}

func TestJson2(t *testing.T) {
	var test = [3]string {"aaa", "bbb", "ccc"}
	test_json, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("test json: %s\n", string(test_json))

	test1_json := string(test_json)
	test1 := make([]string, 0)
	err = json.Unmarshal([]byte(test1_json), &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("test1 slice: %v\n", test1)
}
