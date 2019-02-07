package sort

import "sort"

// 小さいもん順
func Test1(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
}

// 大きいもん順
func Test2(nums []int) {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
}
