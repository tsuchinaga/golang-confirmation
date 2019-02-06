package slice

import "fmt"

// appendしただけでは何の影響もない
func Test1(n int) []int {
	list := make([]int, 0)

	for i := range make([]byte, n) {
		_ = append(list, i)
	}

	return list
}

// appendしたものを上書きし続けたら要素が増える
func Test2(n int) []int {
	list := make([]int, 0)

	for i := range make([]byte, n) {
		list = append(list, i)
	}

	return list
}

// appendを利用して、昇順ソートをかけてみる
func Test3(nums []int) []int {
	list := make([]int, 0)

	for _, n := range nums {
		add := false

		for i, v := range list {
			if n < v {
				add = true
				list = append(list[:i], append([]int{n}, list[i:]...)...)
				break
			}
		}

		if !add {
			list = append(list, n)
		}
	}

	return list
}

// appendを少し変えてみる -> テストが通らない
func Test4(nums []int) []int {
	list := make([]int, 0)

	for _, n := range nums {
		add := false

		for i, v := range list {
			if n < v {
				add = true
				list = append(append(list[:i], n), list[i:]...)
				break
			}
		}

		if !add {
			list = append(list, n)
		}
	}

	return list
}

// i番目で分割したsliceを返す
func Test5(nums []int, i int) ([]int, []int) {
	return nums[:i], nums[i:]
}

// i番目で分割して、後ろのsliceの先頭にnをつけて返す
func Test6(nums []int, i, n int) ([]int, []int) {
	return nums[:i], append([]int{n}, nums[i:]...)
}

// i番目で分割して、前のsliceの最後にnをつけて返す -> 失敗
func Test7(nums []int, i, n int) ([]int, []int) {
	return append(nums[:i], n), nums[i:]
}

// i番目で分割した新しいsliceを作成して、前のsliceの最後にnをつけて返す
func Test8(nums []int, i, n int) ([]int, []int) {
	a := make([]int, 0)
	b := make([]int, 0)

	for j, v := range nums {
		if j < i {
			a = append(a, v)
		} else {
			b = append(b, v)
		}
	}

	return append(a, n), b
}

// 単純に代入では複製されない
func Test9(nums []int, i, n int) ([]int, []int) {
	a := nums
	b := nums

	return append(a[:i], n), b[i:]
}

// copyすることはできる
func Test10(nums []int, i, n int) ([]int, []int) {
	a, b := make([]int, len(nums)), make([]int, len(nums))
	copy(a, nums)
	copy(b, nums)

	return append(a[:i], n), b[i:]
}

// slice自体と、抽出されたsliceのアドレスを見てみる
func Test11(nums []int, i int) (string, string, string) {
	return fmt.Sprintf("%p", nums), fmt.Sprintf("%p", nums[:i]), fmt.Sprintf("%p", nums[i:])
}
