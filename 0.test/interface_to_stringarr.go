package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"strconv"
)

func main() {
	//name:=""
	//
	//for i:=0; i<10;i++  {
	//	name = fmt.Sprintf("%s%d", "abc", i)
	//	println(name)
	//}

	//shellCmd = strings.Replace(shellCmd, "${{CommitID}}", project.CommitID, -1)

	var ii interface{}
	j := "{\"aString\": [\"aaa_111\", \"bbb_222\"], \"whatever\":\"ccc\"}"

	err := json.Unmarshal([]byte(j), &ii)
	if err != nil {
		log.Fatal(err)
	}

	data := ii.(map[string]interface{})
	fmt.Println(data["aString"]) // outputs: ["aaa_111" "bbb_222"]

	var paramSlice []string
	for _, param := range data["aString"].([]interface{}) {
		switch v := param.(type) {
		case string:
			paramSlice = append(paramSlice, v)
		case int:
			strV := strconv.FormatInt(int64(v), 10)
			paramSlice = append(paramSlice, strV)
		default:
			panic("params type not supported")
		}
	}

	println("=======================================")

	for i, v := range paramSlice {
		println(i)
		println(v)
	}

}
