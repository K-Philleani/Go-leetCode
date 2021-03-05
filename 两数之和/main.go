package main

import "fmt"

func twoSum(nums []int, target int) []int {
	 numMap := make(map[int]int)
	 temArr := make([]int, 2)
	 for i, _ := range nums {
	 	 subNum := target - nums[i]
	 	if val, ok := numMap[subNum]; ok {
			temArr[0] = val
			temArr[1] = i
			return temArr
		}
		 numMap[nums[i]] = i
	 }
	 return temArr
}

func main() {
	nums := []int{2,7,11,15}
	target := 9
	res := twoSum(nums, target)
	fmt.Println(res)
}
