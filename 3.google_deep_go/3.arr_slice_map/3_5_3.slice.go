package main

import (
	"fmt"
	"strings"
)

func doubleSliceTest1() [][]string{
	arr1 := make([]string, 0)
	arr1 = append(arr1, "aaa\t")
	arr1 = append(arr1, "bbb\t")

	arr2 := make([]string, 0)
	arr2 = append(arr2, "")
	arr2 = append(arr2, "ddd\t")

	var filteredExcelData [][]string
	filteredExcelData = append(filteredExcelData, arr1)
	filteredExcelData = append(filteredExcelData, arr2)
	//fmt.Println(filteredExcelData)
	return  filteredExcelData
}

func doubleSliceTest2(filteredExcelData1 [][]string) [][]string{
	arr4 := make([]string, 0)
	arr4 = append(arr4, "gg")
	arr4 = append(arr4, "hh")
	filteredExcelData1 = append(filteredExcelData1, arr4)
	return filteredExcelData1
}

func StringFilter(inData string)  string {
	if inData == "" {
		return ""
	}
	return strings.Replace(inData, "\t", "", -1)
}

func main() {
	filteredExcelData := doubleSliceTest1()
	//l1 := len(filteredExcelData)

	arr3 := make([]string, 0)
	arr3 = append(arr3, "eee\t")
	arr3 = append(arr3, "fff\t")
	filteredExcelData = append(filteredExcelData, arr3)
	fmt.Println(filteredExcelData)

	for _, rowData := range filteredExcelData {
		for _, data := range rowData{
			data += "ybx"
		}
	}

	fmt.Println(filteredExcelData)

	/*
	l2 := len(filteredExcelData)
	num := l2-l1
	fmt.Println(num)
	*/

	/*
	var l3 int
	l3=1
	fmt.Println(l3)
	*/

	/*
	filteredExcelData = doubleSliceTest2(filteredExcelData)
	fmt.Println(filteredExcelData)
	*/

}
