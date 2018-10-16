package main

import (
	"fmt"
)

//格式:
//type typeName func(input1 inputType1 , input2 inputType2 [, ...]) (result1 resultType1 [, ...])

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice =", slice)
	odd := filter(slice, isOdd)
	fmt.Println("odd slice =", odd)
	even := filter(slice, isEven)
	fmt.Println("even slice= ", even)
}
