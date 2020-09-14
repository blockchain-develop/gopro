package _interface

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	FUNCTION_MIN_LENGTH = 10
	LETTER = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ*_["
)

type gofunc struct{
	Name string
	Params []string
	ParamTypes []string
	outputs []string
	outputTypes []string
}

func ParseGoFunc(line string) *gofunc {
	gofunc := &gofunc{
	}
	start := 0
	end := 0
	{
		start = strings.IndexAny(line, LETTER)
		end = strings.Index(line[start:], "(")
		end = start + end
		gofunc.Name = line[start : start + end - 1]
		start = end
	}
	for true {
		end = strings.IndexAny(line[start:], LETTER)
		start = start + end
		end = strings.Index(line[start:], " ")
		end = start + end
		gofunc.Params = append(gofunc.Params, line[start:end])

		start = end
		end = strings.IndexAny(line[start:], LETTER)
		start = start + end
		end = strings.IndexAny(line[start:], ",)")
		end = start + end
		gofunc.ParamTypes = append(gofunc.ParamTypes, line[start:end])
		start = end

		if line[end] == ')' {
			break
		}
	}
	end = strings.IndexAny(line[start:], "(")
	if end == -1 {
		return gofunc
	}
	start = start + end
	for true {
		end = strings.IndexAny(line[start:], LETTER)
		start = start + end
		end = strings.IndexAny(line[start:], ",)")
		end = start + end
		gofunc.outputs = append(gofunc.outputs, "")
		gofunc.outputTypes = append(gofunc.outputTypes, line[start:end])
		start = end

		if line[end] == ')' {
			break
		}
	}
	return gofunc
}

func MappingTypes(in string) string {
	if in == "common.Address" {
		return "address"
	} else if in == "*big.Int" {
		return "uint256"
	} else if in == "big.Int" {
		return "uint256"
	}
	return in
}

func ToJson(gofunc *gofunc) string {
	value := "{\"type\":\"function\",\"constant\":true,\"name\":\"` + "
	value += gofunc.Name
	value += " + `\",\"inputs\":["
	for i, param := range gofunc.Params {
		value += "{\"name\":\""
		value += param
		value += "\",\"type\":\""
		value += MappingTypes(gofunc.ParamTypes[i])
		value += "\"}"
		if i != len(gofunc.Params) - 1 {
			value += ","
		}
	}
	value += "],\"outputs\":["
	for i, output := range gofunc.outputs {
		value += "{\"name\":\""
		value += output
		value += "\",\"type\":\""
		value += MappingTypes(gofunc.outputTypes[i])
		value += "\"}"
		if i != len(gofunc.outputs) - 1 {
			value += ","
		}
	}
	value += "],\"payable\":false,\"stateMutability\":\"nonpayable\"}"
	return value
}

func TestInterfaceParser(t *testing.T) {
	f, err := os.Open("./erc721.go")
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	content := string(data)
	lines := strings.Split(content, "\n")


	command := false
	ininterface := false
	gofuncs := make([]*gofunc, 0)
	for _, line := range lines {
		if command {
			if strings.Contains(line, "*/") {
				command = false
			}
			continue
		}
		if strings.Contains(line, "/*") {
			command = true
			continue
		}
		if strings.Contains(line, "//") {
			continue
		}
		if !ininterface {
			if strings.Contains(line, "interface {") {
				ininterface = true
			}
			continue
		}
		if strings.Contains(line, "}") && len(line) < FUNCTION_MIN_LENGTH {
			ininterface = false
		}
		if strings.Contains(line, "(") && strings.Contains(line, ")") {
			gofunc := ParseGoFunc(line)
			gofuncs = append(gofuncs, gofunc)
		}
	}

	outputJson := "["
	for i, gofunc := range  gofuncs {
		outputJson += ToJson(gofunc)
		if i != len(gofuncs) - 1 {
			outputJson += ",\n"
		}
	}
	outputJson += "]"
	fmt.Printf("output json: \n%s\n",outputJson)
}
