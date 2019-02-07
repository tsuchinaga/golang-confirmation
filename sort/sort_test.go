package sort

import (
	"reflect"
	"testing"
)

func TestTest1(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{3, 3, 1}, []int{1, 3, 3}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for i, test := range testTable {
		Test1(test.nums)

		if !reflect.DeepEqual(test.expect, test.nums) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, test.nums)
		}
	}
}

func TestTest2(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 1}, []int{3, 2, 1}},
		{[]int{3, 3, 1}, []int{3, 3, 1}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
	}

	for i, test := range testTable {
		Test2(test.nums)

		if !reflect.DeepEqual(test.expect, test.nums) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, test.nums)
		}
	}
}
