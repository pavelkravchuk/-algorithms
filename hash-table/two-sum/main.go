package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	usedNums := make(map[int]int)

	for pos, num := range nums {

		if usedNum, exist := usedNums[target-num]; exist {
			result = append(result, usedNum)
			result = append(result, pos)
			return result
		}
		usedNums[num] = pos
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}
