package sort

import (
	"sort"
)

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

// 手動でソートしたときの入れ替え回数 ※与えられる数字はすべて違い、1からの連番であることが決まっている場合
func Test3(nums []int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for {
			if i+1 == nums[i] {
				break
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			cnt++
		}
	}
	return cnt
}

// sort.Sliceでの入れ替え回数 ※与えられる数字はすべて違い、1からの連番であることが決まっている場合
func Test4(nums []int) int {
	cnt := 0
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < nums[j] {
			cnt++
			return true
		}
		return false
	})
	return cnt
}
