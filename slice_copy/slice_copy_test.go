package slice_copy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceCopy1(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		expect1 []int
		expect2 []int
	}{
		{[]int{1, 2, 3}, 0, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1}, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}, []int{3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := SliceCopy1(test.nums, test.i)

		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}

		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}

		if fmt.Sprintf("%p", test.nums) == fmt.Sprintf("%p", actual1) || fmt.Sprintf("%p", test.nums) == fmt.Sprintf("%p", actual2) {
			t.Fatalf("%s No.%02d 失敗\n%p %p %p", t.Name(), i+1, test.nums, actual1, actual2)
		}
	}
}

func TestSliceCopy2(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		expect1 []int
		expect2 []int
	}{
		{[]int{1, 2, 3}, 0, []int{}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 1, []int{1}, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}, []int{3}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}, []int{}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := SliceCopy2(test.nums, test.i)

		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}

		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}

		if fmt.Sprintf("%p", test.nums) == fmt.Sprintf("%p", actual1) || fmt.Sprintf("%p", test.nums) == fmt.Sprintf("%p", actual2) {
			t.Fatalf("%s No.%02d 失敗\n%p %p %p", t.Name(), i+1, test.nums, actual1, actual2)
		}
	}
}
