package error

import (
	"errors"
	"fmt"
	"testing"
)

func TestWrapError(t *testing.T) {
	testTable := []struct {
		desc    string
		message string
	}{
		{desc: "messageが空", message: ""},
		{desc: "messageが空じゃない", message: "foo bar"},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			fmt.Printf("%s -> %v\n", t.Name(), errors.Unwrap(WrapErrorFromBase(test.message)))
			fmt.Printf("%s -> %v\n", t.Name(), errors.As(WrapErrorFromBase(test.message), &BaseError))
			fmt.Printf("%s -> %v\n", t.Name(), errors.Is(WrapErrorFromBase(test.message), BaseError))
		})
	}
}

func TestWrapErrorFromSuper(t *testing.T) {
	testTable := []struct {
		desc    string
		message string
	}{
		{desc: "messageが空"},
		{desc: "messageが空じゃない", message: "foo bar baz"},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
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

func TestErrorsIs(t *testing.T) {
	t.Log("BaseError", "BaseError", errors.Is(BaseError, BaseError))
	t.Log("SuperError", "SuperError", errors.Is(SuperError, SuperError))
	t.Log("BaseError", "SuperError", errors.Is(BaseError, SuperError))
	t.Log("SuperError", "BaseError", errors.Is(SuperError, BaseError))
}

func TestNilErrorIs(t *testing.T) {
	var nilError error
	wrappedNilError := fmt.Errorf("%w: wraped", nilError)
	t.Log(errors.Is(nilError, nil))
	t.Log(errors.Is(nil, nilError))

	t.Log(errors.Is(nilError, wrappedNilError))
	t.Log(errors.Is(wrappedNilError, nilError))
	t.Log(errors.Is(nil, wrappedNilError))
	t.Log(errors.Is(wrappedNilError, nil))

	t.Log(errors.Is(errors.Unwrap(wrappedNilError), nil))
	t.Log(errors.Is(nil, errors.Unwrap(wrappedNilError)))
}
