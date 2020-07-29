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

}
