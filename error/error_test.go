package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	testTable := []struct {
		dist    string
		message string
	}{
		{dist: "messageが空", message: ""},
		{dist: "messageが空じゃない", message: "foo bar"},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.dist, func(t *testing.T) {
			t.Parallel()

			fmt.Printf("%s -> %v\n", t.Name(), errors.Unwrap(WrapErrorFromBase(test.message)))
			fmt.Printf("%s -> %v\n", t.Name(), errors.As(WrapErrorFromBase(test.message), &BaseError))
			fmt.Printf("%s -> %v\n", t.Name(), errors.Is(WrapErrorFromBase(test.message), BaseError))
		})
	}
}

func TestWrapErrorFromSuper(t *testing.T) {
	testTable := []struct {
		dist    string
		message string
	}{
		{dist: "messageが空"},
		{dist: "messageが空じゃない", message: "foo bar baz"},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.dist, func(t *testing.T) {
			// t.Parallel()
			fmt.Printf("%s -> %v\n", t.Name(), errors.Unwrap(WrapErrorFromSuper(test.message)))
			fmt.Printf("%s -> %v\n", t.Name(), errors.Unwrap(errors.Unwrap(WrapErrorFromSuper(test.message))))
			fmt.Printf("%s -> %v\n", t.Name(), errors.As(WrapErrorFromSuper(test.message), &BaseError))
			fmt.Printf("%s -> %v\n", t.Name(), errors.As(WrapErrorFromSuper(test.message), &SuperError))
			fmt.Printf("%s -> %v\n", t.Name(), errors.Is(WrapErrorFromSuper(test.message), BaseError))
			fmt.Printf("%s -> %v\n", t.Name(), errors.Is(WrapErrorFromSuper(test.message), SuperError))
		})
	}
}
