package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

type Contract struct {
	Address         string  `json:"address"`
	Chain           uint32  `json:"chain"`
	IsOK            bool
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

/*
将json Unmarshal为结构体对象
json: {"address":"aaa","chain":0}
*/
func TestJson_Unmarshal1(t *testing.T) {
	var test = &Contract {
		Address: "aaa",
		Chain:   0,
		IsOK:false,
	}
	test_json, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	var test1 Contract
	err = json.Unmarshal(test_json, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", test1)
}

/*
将json Unmarshal为slice
json: ["aaa","bbb","ccc"]
*/
func TestJson_Unmarshal2(t *testing.T) {
	var test = [3]string {"aaa", "bbb", "ccc"}
	test_json, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	test1_json := string(test_json)
	test1 := make([]string, 0)
	err = json.Unmarshal([]byte(test1_json), &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value: %v\n", test1)
}

/*
将json Unmarshal为结构体对象的slice
json: [{"address":"aaa","chain":0},{"address":"bbb","chain":0},{"address":"ccc","chain":0}]
*/
func TestJson_Unmarshal3(t *testing.T) {
	var test = [3]*Contract {
		{
			Address: "aaa",
			Chain:   0,
		},
		{
			Address: "bbb",
			Chain:   0,
		},
		{
			Address: "ccc",
			Chain:   0,
		},
	}
	test_json, err := json.Marshal(test)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	test1 := make([]*Contract, 0)
	err = json.Unmarshal(test_json, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", test1)
}

/*
将json Unmarshal为map
json: {"contract1":{"name":"aaa","address":{"aaa":"aaa","bbb":"bbb"}},"contract2":{"name":"bbb","address":{"aaa":"aaa","bbb":"bbb"}}}
*/
func TestJson_Unmarshal4(t *testing.T) {
	type Contract1 struct {
		Name   string  `json:"name"`
		Address  map[string]string  `json:"address"`
	}
	address := make(map[string]string, 0)
	address["aaa"] = "aaa"
	address["bbb"] = "bbb"
	contract1 := &Contract1{
		Name:    "aaa",
		Address: address,
	}
	contract2 := &Contract1{
		Name:    "bbb",
		Address: address,
	}
	data := make(map[string]*Contract1)
	data["contract1"] = contract1
	data["contract2"] = contract2

	test_json, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	test1 := make(map[string]*Contract1, 0)
	err = json.Unmarshal(test_json, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", test1)
}

/*
默认的Marshal
{"Name":"aaa","Address":"bbb"}
*/
func TestJson_Marshal1(t *testing.T) {
	type Contract1 struct {
		Name   string
		Address  string
	}
	data := &Contract1{
		Name:    "aaa",
		Address: "bbb",
	}
	test_json, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))
}

/*
指定别名的Marshal
{"name":"aaa","location":"bbb"}
*/

func TestJson_Marshal2(t *testing.T) {
	type Contract1 struct {
		Name   string  `json:"name"`
		Address  string  `json:"location"`
	}
	data := &Contract1{
		Name:    "aaa",
		Address: "bbb",
	}
	test_json, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))
}


/*
忽略字段的Marshal
{"name":"aaa"}
*/

func TestJson_Marshal3(t *testing.T) {
	type Contract1 struct {
		Name   string  `json:"name"`
		Address  string  `json:"-"`  /*忽略字段*/
		Card     string   `json:",omitempty"`  /*如果内容为空，则忽略字段*/
	}
	data := &Contract1{
		Name:    "aaa",
		Address: "bbb",
	}
	test_json, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))
}


/*
类型转换的Marshal
{"name":"aaa","age":"1","amount":"0.001"}
*/
func TestJson_Marshal4(t *testing.T) {
	type Contract1 struct {
		Name   string      `json:"name"`
		Age     int        `json:"age,string"`
		Amount  float64    `json:"amount,string"`
	}
	data := &Contract1{
		Name:    "aaa",
		Age: 	1,
		Amount: float64(0.001),
	}
	test_json, err := json.Marshal(&data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	var data1 Contract1
	err = json.Unmarshal(test_json, &data1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", data1)
}

/*
自定义类型的json Marshal
*/
type BigInt struct {
	big.Int
}

func (bigInt *BigInt) MarshalJSON() ([]byte, error) {
	if  bigInt == nil {
		return []byte("null"), nil
	}
	return []byte(bigInt.String()), nil
}

func (bigInt *BigInt) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	data, ok := new(big.Int).SetString(string(p), 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	bigInt.Int = *data
	return nil
}

func TestJson_Marshal5(t *testing.T) {
	type Contract2 struct {
		Name   string      `json:"name"`
		Age    int         `json:"age,string"`
		Amount  float64    `json:"amount,string"`
		Asset *BigInt `json:"asset"`
	}
	data, _ := new(big.Int).SetString("111111111111111111111111111111111111111111111111111111", 10)
	contract := &Contract2{
		Name:   "aaaa",
		Age:    12,
		Amount: float64(0.2345),
		Asset:  &BigInt{*data,},
	}
	test_json, err := json.Marshal(&contract)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Marshal json: %s\n", string(test_json))

	var data1 Contract2
	err = json.Unmarshal(test_json, &data1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", data1)
}



