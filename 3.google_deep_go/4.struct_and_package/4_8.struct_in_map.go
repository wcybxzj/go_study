package main

import (
	"encoding/json"
	"fmt"
)

func LogDataGrpc( data interface{})  {
	var dataStr string
	bs, err := json.Marshal(data)
	if err == nil {
		dataStr = string(bs)
	} else {
		dataStr = ""
	}
	fmt.Println(dataStr)
}

func test1_struct_in_map() {
	type Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Addr string `json:"addr"`
	}

	m := make(map[string]interface{})
	m["name"] = "ybx"
	m["st"] = Person{
		Name: "taoge",
		Age:  30,
		Addr: "China",  // oh my god, this comma cannot be omitted
	}

	LogDataGrpc(m)
}

func main() {
	test1_struct_in_map()
}
