package append

import (
	"reflect"
	"testing"
)

func TestTest1(t *testing.T) {
	testTable := []struct {
		n      int
		expect []int
	}{
		{0, make([]int, 0)},
		{1, make([]int, 0)},
		{3, make([]int, 0)},
		{5, make([]int, 0)},
	}

	for i, test := range testTable {
		actual := Test1(test.n)
		if !reflect.DeepEqual(test.expect, actual) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestTest2(t *testing.T) {
	testTable := []struct {
		n      int
		expect []int
	}{
		{0, []int{}},
		{1, []int{0}},
		{3, []int{0, 1, 2}},
		{5, []int{0, 1, 2, 3, 4}},
	}

	for i, test := range testTable {
		actual := Test2(test.n)
		if !reflect.DeepEqual(test.expect, actual) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestTest3(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{2, 1, 0}, []int{0, 1, 2}},
		{[]int{1, 1, 0}, []int{0, 1, 1}},
		{[]int{5, 1, 4, 2, 3}, []int{1, 2, 3, 4, 5}},
	}

	for i, test := range testTable {
		actual := Test3(test.nums)
		if !reflect.DeepEqual(test.expect, actual) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestTest4(t *testing.T) {
	testTable := []struct {
		nums   []int
		expect []int
	}{
		{[]int{}, []int{}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{2, 1, 0}, []int{0, 1, 2}},
		{[]int{1, 1, 0}, []int{0, 1, 1}},
		{[]int{5, 1, 4, 2, 3}, []int{1, 2, 3, 4, 5}},
	}

	for i, test := range testTable {
		actual := Test4(test.nums)
		if !reflect.DeepEqual(test.expect, actual) {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestTest5(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, []int{}, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, []int{0}, []int{1, 2}},
		{[]int{0, 1, 2}, 2, []int{0, 1}, []int{2}},
		{[]int{0, 1, 2}, 3, []int{0, 1, 2}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test5(test.nums, test.i)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest6(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		n       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, 5, []int{}, []int{5, 0, 1, 2}},
		{[]int{0, 1, 2}, 1, 5, []int{0}, []int{5, 1, 2}},
		{[]int{0, 1, 2}, 2, 5, []int{0, 1}, []int{5, 2}},
		{[]int{0, 1, 2}, 3, 5, []int{0, 1, 2}, []int{5}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test6(test.nums, test.i, test.n)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest7(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		n       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, 5, []int{5}, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, 5, []int{0, 5}, []int{1, 2}},
		{[]int{0, 1, 2}, 2, 5, []int{0, 1, 5}, []int{2}},
		{[]int{0, 1, 2}, 3, 5, []int{0, 1, 2, 5}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test7(test.nums, test.i, test.n)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest8(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		n       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, 5, []int{5}, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, 5, []int{0, 5}, []int{1, 2}},
		{[]int{0, 1, 2}, 2, 5, []int{0, 1, 5}, []int{2}},
		{[]int{0, 1, 2}, 3, 5, []int{0, 1, 2, 5}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test8(test.nums, test.i, test.n)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest9(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		n       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, 5, []int{5}, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, 5, []int{0, 5}, []int{1, 2}},
		{[]int{0, 1, 2}, 2, 5, []int{0, 1, 5}, []int{2}},
		{[]int{0, 1, 2}, 3, 5, []int{0, 1, 2, 5}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test9(test.nums, test.i, test.n)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest10(t *testing.T) {
	testTable := []struct {
		nums    []int
		i       int
		n       int
		expect1 []int
		expect2 []int
	}{
		{[]int{0, 1, 2}, 0, 5, []int{5}, []int{0, 1, 2}},
		{[]int{0, 1, 2}, 1, 5, []int{0, 5}, []int{1, 2}},
		{[]int{0, 1, 2}, 2, 5, []int{0, 1, 5}, []int{2}},
		{[]int{0, 1, 2}, 3, 5, []int{0, 1, 2, 5}, []int{}},
	}

	for i, test := range testTable {
		actual1, actual2 := Test10(test.nums, test.i, test.n)
		if !reflect.DeepEqual(test.expect1, actual1) {
			t.Fatalf("%s No.%02d-1 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect1, actual1)
		}
		if !reflect.DeepEqual(test.expect2, actual2) {
			t.Fatalf("%s No.%02d-2 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect2, actual2)
		}
	}
}

func TestTest11(t *testing.T) {
	nums := []int{0, 1, 2}

	testTable := []struct {
		nums []int
		i    int
	}{
		{nums, 0},
		{nums, 1},
		{nums, 2},
		{nums, 3},
	}

	for i, test := range testTable {
		actual1, actual2, actual3 := Test11(test.nums, test.i)
		t.Logf("%s No.%02d: actual1=%s, actual2=%s, actual3=%s", t.Name(), i+1, actual1, actual2, actual3)
	}
}
