package main

import (
	"reflect"
	"testing"
	"time"
)

func TestGetMonth(t *testing.T) {
	testTable := []struct {
		i      int
		expect string
	}{
		{1, "January"},
		{2, "February"},
		{3, "March"},
		{4, "April"},
		{5, "May"},
		{6, "June"},
		{7, "July"},
		{8, "August"},
		{9, "September"},
		{10, "October"},
		{11, "November"},
		{12, "December"},
		{0, "%!Month(0)"},
		{13, "%!Month(13)"},
	}
	for i, test := range testTable {
		actual := GetMonth(test.i)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v(%+v)\n実際: %+v(%+v)\n", t.Name(), i+1,
				test.expect, reflect.TypeOf(test.expect), actual, reflect.TypeOf(actual))
		}
	}
}

func TestEqual(t *testing.T) {
	testTable := []struct {
		a, b   time.Time
		expect bool
	}{
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), true},
		{time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), false},
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), false},
	}
	for i, test := range testTable {
		actual := Equal(test.a, test.b)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestGt(t *testing.T) {
	testTable := []struct {
		a, b   time.Time
		expect bool
	}{
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), false},
		{time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), true},
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), false},
	}
	for i, test := range testTable {
		actual := Gt(test.a, test.b)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestGe(t *testing.T) {
	testTable := []struct {
		a, b   time.Time
		expect bool
	}{
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), true},
		{time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), true},
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), false},
	}
	for i, test := range testTable {
		actual := Ge(test.a, test.b)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestLt(t *testing.T) {
	testTable := []struct {
		a, b   time.Time
		expect bool
	}{
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), false},
		{time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), false},
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), true},
	}
	for i, test := range testTable {
		actual := Lt(test.a, test.b)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestLe(t *testing.T) {
	testTable := []struct {
		a, b   time.Time
		expect bool
	}{
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), true},
		{time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), false},
		{time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local), time.Date(2010, 1, 1, 0, 0, 0, 1, time.Local), true},
	}
	for i, test := range testTable {
		actual := Le(test.a, test.b)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestBetween(t *testing.T) {
	testTable := []struct {
		a, b, c time.Time
		expect  bool
	}{
		{
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			true,
		}, {
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2011, 1, 1, 0, 0, 0, 0, time.Local),
			true,
		}, {
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2011, 1, 1, 0, 0, 0, 0, time.Local),
			true,
		}, {
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			true,
		}, {
			time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			false,
		}, {
			time.Date(2011, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			false,
		}, {
			time.Date(2010, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2011, 1, 1, 0, 0, 0, 0, time.Local),
			time.Date(2009, 1, 1, 0, 0, 0, 0, time.Local),
			false,
		},
	}
	for i, test := range testTable {
		actual := Between(test.a, test.b, test.c)
		if test.expect != actual {
			t.Fatalf("%s No.%02d 失敗\n期待: %+v\n実際: %+v\n", t.Name(), i+1, test.expect, actual)
		}
	}
}
