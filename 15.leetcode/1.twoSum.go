package main

import "fmt"

//https://leetcode.com/problems/two-sum/
//https://leetcode-cn.com/problems/two-sum/
//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。


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


//方法2:two-way-hash
func twoSumV2(nums []int, target int) []int {
	hash1_map := make(map[int]int)
	//hash2_map := make(map[int]int)

	for k, v :=range nums {
		hash1_map[v] = k
	}

	for k, v :=range nums {
		tmp := target-v
		k1, ok := hash1_map[tmp]
		if ok && k != k1{
			return []int{k ,k1}
		}
	}

	return []int{}
}



//方法3:one-way-hash
func twoSumV3(nums []int, target int) []int {
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


func main() {
	var ret []int
	arr1 := []int{2, 7, 11, 15}

	ret = twoSumV1(arr1, 9)
	fmt.Println(ret)

	ret = twoSumV2(arr1, 9)
	fmt.Println(ret)

	ret = twoSumV3(arr1, 9)
	fmt.Println(ret)
}
