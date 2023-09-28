package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func Test_json_marshal(t *testing.T) {
	var p = func(b []byte) {
		fmt.Println(string(b))
	}

	boolByte, _ := json.Marshal(true)
	p(boolByte)

	intByte, _ := json.Marshal(1)
	p(intByte)

	floatByte, _ := json.Marshal(1.23)
	p(floatByte)

	strByte, _ := json.Marshal("str")
	p(strByte)

	strSliceByte, _ := json.Marshal([]string{"a", "b", "c"})
	p(strSliceByte)

	mapByte, _ := json.Marshal(map[string]int{"A": 1, "B": 2})
	p(mapByte)

	type s struct {
		a1 int
		A1 int
		A2 int `json:"AX"`
	}
	structByte, _ := json.Marshal(s{
		a1: 1, // only Upper-word-filed are exportable
		A1: 1,
		A2: 1,
	})
	p(structByte)

	encoder := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	encoder.Encode(d)
}

func Test_json_unmarshal(t *testing.T) {
	dataByte := []byte(`
		{
			"A1": 2,
			"AX": 2
		}
	`)
	var data map[string]interface{}
	err := json.Unmarshal(dataByte, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	fmt.Println(data["AX"])

	var x struct {
		A1 int
		A2 int `json:"AX"`
	}
	err = json.Unmarshal(dataByte, &x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#+v\n", x)
}
