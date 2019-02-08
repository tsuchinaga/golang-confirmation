package slice_copy

// numsをi番目で切った配列を、copyして返す
func SliceCopy1(nums []int, i int) ([]int, []int) {

	// out of length回避
	if i < 0 || len(nums)-1 < i {
		i = len(nums)
	}

	a, b := make([]int, i), make([]int, len(nums)-i)
	copy(a, nums[:i])
	copy(b, nums[i:])
	return a, b
}

// numsをi番目で切った配列を、copyして返す ※aへのcopy方法を少しかえただけ
func SliceCopy2(nums []int, i int) ([]int, []int) {

	// out of length回避
	if i < 0 || len(nums)-1 < i {
		i = len(nums)
	}

	a, b := make([]int, i), make([]int, len(nums)-i)
	copy(a, nums)
	copy(b, nums[i:])
	return a, b
}
