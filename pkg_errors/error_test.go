package error

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
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

func Test_WrapOriginalErrorFromBase(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		is   error
		want bool
	}{
		{name: "original errorをwrapしていない", is: OriginalError, want: false},
		{name: "base errorをwrapしている", is: BaseError, want: true},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := WrapOriginalErrorFromBase()
			if errors.Is(got, test.is) != test.want {
				t.Errorf("%s error\nwant: %+v\ngot: %+v, unwrap: %+v\n", t.Name(), test.is, got, errors.Unwrap(got))
			} else {
				t.Logf("%+v\n", got)
			}
		})
	}
}

func Test_WrapBaseErrorFromOriginal(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		is   error
		want bool
	}{
		{name: "original errorをwrapしている", is: OriginalError, want: true},
		{name: "base errorをwrapしていない", is: BaseError, want: false},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			got := WrapBaseErrorFromOriginal()
			if errors.Is(got, test.is) != test.want {
				t.Errorf("%s error\nwant: %+v\ngot: %+v, unwrap: %+v\n", t.Name(), test.is, got, errors.Unwrap(got))
			} else {
				t.Logf("%+v\n", got)
			}
		})
	}
}
