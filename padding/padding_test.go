package padding

import "testing"

func TestPadding1(t *testing.T) {
	testTable := []struct {
		str    string
		expect string
	}{
		{"", "        "},
		{"    ", "        "},
		{"a", "       a"},
		{"abcdef", "  abcdef"},
		{"abcdefghhijk", "abcdefghhijk"},
	}

	for i, test := range testTable {
		actual := Padding1(test.str)
		if test.expect != actual {
			t.Fatalf("%s No.02%d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}

func TestPadding2(t *testing.T) {
	testTable := []struct {
		str    string
		expect string
	}{
		{"", "        "},
		{"    ", "        "},
		{"a", "a       "},
		{"abcdef", "abcdef  "},
		{"abcdefghhijk", "abcdefghhijk"},
	}

	for i, test := range testTable {
		actual := Padding2(test.str)
		if test.expect != actual {
			t.Fatalf("%s No.02%d 失敗\n期待: %+v\n実際: %+v", t.Name(), i+1, test.expect, actual)
		}
	}
}
