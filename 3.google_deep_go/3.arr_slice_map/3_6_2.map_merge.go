package main

import "fmt"

func mapMerge(m1 map[string]interface{}, m2 map[string]interface{}, cover bool) map[string]interface{} {
	for index, data := range m2 {
		_, ok := m1[index]
		if ok {
			if cover == false {
				continue
			}
		}
		m1[index] = data
	}
	return m1
}

func main() {
	m1 := make(map[string]interface{})
	m1["name"] = "ybx"
	m1["age"] = 123

	m2 := make(map[string]interface{})
	m2["name"]="wc"

	m3 := make(map[string]interface{})
	m3["name"]="ly"
	m3["sex"]="woman"

	m := mapMerge(m1 ,m2, true)
	fmt.Println(m)

	m = mapMerge(m2 ,m3, false)
	fmt.Println(m)
}

