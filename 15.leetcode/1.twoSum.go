package main

import "fmt"

//方法1:暴力循环
func twoSumV1(nums []int, target int) []int {
	var ret []int
	len := len(nums)
	for i := 0; i < len; i++ {
		for j := i + 1; j < len; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
		fmt.Println(nums[i])
	}
	return ret
}

//方法2:hash
func twoSumV2(nums []int, target int) []int {
	var ret []int
	hash_map := make(map[int]int)
	for i, x := range nums {
		_, ok := hash_map[target-x]
		if ok {
			ret = append(ret, hash_map[target-x])
			ret = append(ret, i)
			return ret
		}
		hash_map[x] = i
	}

	return ret
}

//https://leetcode.com/problems/two-sum/
func main() {
	var ret []int
	arr1 := []int{2, 7, 11, 15}

	ret = twoSumV1(arr1, 9)
	fmt.Println(ret)

	ret = twoSumV2(arr1, 9)
	fmt.Println(ret)
}
