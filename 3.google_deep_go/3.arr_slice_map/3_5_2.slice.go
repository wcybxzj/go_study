package main

import (
	"fmt"
	"zuji/common/math"
)

func CompareTwoStringSlices(s1[]string, s2[]string) []string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func GetMaxLenSlice(slices ...[]string) []string {
	if len(slices) < 2 {
		return slices[0]
	}
	slice := slices[0]
	for i := 1; i < len(slices); i++ {
		slice = CompareTwoStringSlices(slice, slices[i])
	}
	return slice
}

func test1(s1[]string, s2[]string, s3[]string)([][]string) {
	s1Len := len(s1)
	s2Len := len(s2)
	s3Len := len(s3)

	maxLen := math.Max(int64(s1Len), int64(s2Len), int64(s3Len))
	numInt32 := int(maxLen)

	data := make([][]string, maxLen)
	for i:=0; i < numInt32; i++ {
		data[i] = make([]string, 3)
		if i < s1Len {
			data[i][0] = s1[i]
		}

		if i < s2Len {
			data[i][1] = s2[i]
		}

		if i < s3Len {
			data[i][2] = s3[i]
		}
	}
	return data
}

func unit_test1() {
	s1 := make([]string, 0)
	s1 = append(s1, "a1")
	s1 = append(s1, "a2")

	s2 := make([]string, 0)
	s2 = append(s2, "b1")
	s2 = append(s2, "b2")
	s2 = append(s2, "b3")

	s3 := make([]string, 0)

	doubleData := test1(s1, s2, s3)

	for _, data :=range doubleData {
		for i:=0; i<len(data);i++ {
			if data[i]!="" {
				fmt.Print(data[i])
			}else{
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func test2(invalidFormatMobiles []string) (arr1[]string) {
	return []string{"aaa","bbb"}
	//return arr1
}

func test3()  {
	var s1 = [...][]string{{"111", "222"},{"333"},{"4444", "555","6666"},{""}}
	fmt.Println(s1)
}

func test4()  {
	arr1 := make([]string, 0)
	arr1 = append(arr1, "123")


	if len(arr1) == 1 {
		val := arr1[0]

		if val =="" {
			fmt.Println("not set")
		}
	}
}

func main() {
	str1 := "abc123"
	if len(str1)>=2 {
		begin := str1[0:2]
		if begin == "1ab" {
			fmt.Println("ab beign")
		}else{
			fmt.Println("not ab beign")
		}
	}
}
