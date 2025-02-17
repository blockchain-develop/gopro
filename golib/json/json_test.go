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

func TestJson_Unmarshal_empty(t *testing.T) {
	var test1 Contract
	//err := json.Unmarshal([]byte(""), &test1)
	err := json.Unmarshal([]byte{}, &test1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Unmarshal value : %v\n", test1)
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

type BBB struct {
	CCC string
}

type DDD struct {
	t int
	BBB BBB
	CCC string
}

func (d *DDD) UnmarshalJSON(data []byte) error {
	if data[0] == '{' {
		d.t = 0
		return json.Unmarshal(data, &d.BBB)
	} else {
		d.t = 1
		d.CCC = string(data)
		return nil
	}
}

type Data struct {
	BBBs []DDD
}

func TestJson_Unmashall1(t *testing.T) {
	var data Data
	err := json.Unmarshal([]byte(`{"BBBs":[{"CCC":"1111"},"2222"]}`), &data)
	if err != nil {
		panic(err)
	}
}

/*
func TestJson_Unmashall(t *testing.T) {
	data := Data{
		BBBs: make([]BBB, 0),
	}
	data.BBBs = append(data.BBBs, BBB{CCC:"1111"})
	data.BBBs = append(data.BBBs, BBB{CCC:"2222"})
	xx, _ := json.Marshal(data)
	fmt.Printf("xx: %s\n", string(xx))
}

 */


type XXX struct {
	Chain    string `json:"chain"`
	Symbol   string `json:"symbol"`
	Contract string `json:"contract"`
	Decimal  string `json:"decimal"`
	Denom    string `json:"denom"`
}

func TestJson_chainwallet(t *testing.T) {
	fileName := "./test2.json"
	data, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}

	aaa :=
		struct {
			Data []*XXX `json:"data"`
			Code uint `json:"code"`
		}{}

	//
	err = json.Unmarshal(data, &aaa)
	if err != nil {
		fmt.Printf("json.Unmarshal TestConfig:%s error:%s", data, err)
	}

	//
	bbb := make(map[string]*XXX)
	for _, item := range aaa.Data {
		bbb[item.Chain + ":" + item.Symbol] = item
	}
	dataJson, _ := json.MarshalIndent(bbb, "", "    ")
	fmt.Printf("data: \n%s\n", dataJson)
}

type Result struct {
	Cardano Cardano `json:"cardano"`
}

type Cardano struct {
	Outputs []*Output `json:"outputs"`
}

type Output struct {
	Transaction struct {
		Hash string `json:"hash"`
	} `json:"transaction"`
	Currency struct {
		Name string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currency"`
	OutputIndex uint `json:"outputIndex"`
}

type Unspend struct {
	Coin string `json:"Coin"`
	Txid  string `json:"Txid"`
	Index uint `json:"N"`
}

func TestJson_cardano(t *testing.T) {
	fileName := "./cardano1.json"

	data1, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}
	var result1 Result
	json.Unmarshal(data1, &result1)

	fileName = "./cardano2.json"
	data2, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}
	var result2 Result
	json.Unmarshal(data2, &result2)

	fileName = "./cardano3.json"
	data3, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}
	var result3 Result
	json.Unmarshal(data3, &result3)

	fileName = "./cardano_unspent.json"
	data4, err := readFile(fileName)
	if err != nil {
		fmt.Printf("read file err: %v\n", err)
	}

	//
	result1_1 := make(map[string][]string)
	for _, item := range result1.Cardano.Outputs {
		key := fmt.Sprintf("%s:%d", item.Transaction.Hash, item.OutputIndex)
		value, ok := result1_1[key]
		if !ok {
			value = make([]string, 0)
			result1_1[key] = value
		}
		value = append(value, item.Currency.Symbol)
		result1_1[key] = value
	}

	for _, item := range result2.Cardano.Outputs {
		key := fmt.Sprintf("%s:%d", item.Transaction.Hash, item.OutputIndex)
		value, ok := result1_1[key]
		if !ok {
			value = make([]string, 0)
			result1_1[key] = value
		}
		value = append(value, item.Currency.Symbol)
		result1_1[key] = value
	}

	for _, item := range result3.Cardano.Outputs {
		key := fmt.Sprintf("%s:%d", item.Transaction.Hash, item.OutputIndex)
		value, ok := result1_1[key]
		if !ok {
			value = make([]string, 0)
			result1_1[key] = value
		}
		value = append(value, item.Currency.Symbol)
		result1_1[key] = value
	}

	result4 := make([]*Unspend, 0)
	json.Unmarshal(data4, &result4)

	result2_1 := make(map[string][]string)
	for _, item := range result4 {
		key := fmt.Sprintf("%s:%d", item.Txid, item.Index)
		value, ok := result2_1[key]
		if !ok {
			value = make([]string, 0)
			result2_1[key] = value
		}
		value = append(value, item.Coin)
		result2_1[key] = value
	}

	//
	for k1, v1 := range result2_1 {
		v2, ok := result1_1[k1]
		if !ok {
			fmt.Printf("%s is not exist\n", k1)
			continue
		}
		fmt.Printf("%s: %s, %s\n", k1, v1, v2)
	}

}


