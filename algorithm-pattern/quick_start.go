package main

import "fmt"

func main() {
	_ = strStr("this is stack", "ck")
	nums := make([]int, 0)
	nums = append(nums, 1, 2, 3, 4, 5, 6, 7)
	_ = subsets(nums)
	subsetsMy(nums)
}

// strStr 28  196 30.30%
func strStr(haystack, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	var i, j int

	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if len(needle) == j {
			return i
		}
	}

	return -1
}

func subsetsMy(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	for _, num := range nums {
		list = append(list, num)
		for _, listItem := range list {
			list = append(list, listItem)
			result = append(result, list)
		}
	}
	fmt.Println(len(result))
	fmt.Println(len(list))
	return [][]int{}
}

// 78 子集 0ms  100%
func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

func backtrack(nums []int, pos int, list []int, result *[][]int) {
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
