package copy

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	t.Parallel()

	s := "string"
	actual := String(s)

	s = "string2"
	if s == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), s, actual)
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	n := 3
	actual := Int(n)

	n++
	if n == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), n, actual)
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	f := 3.5
	actual := Float(f)

	f++
	if f == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), f, actual)
	}
}

func TestAge(t *testing.T) {
	t.Parallel()

	a := AgeStruct(28)
	actual := Age(a)

	a++
	if a == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), a, actual)
	}
}

func TestAgeStruct_Copy(t *testing.T) {
	t.Parallel()

	a := AgeStruct(28)
	actual := a.Copy()

	a++
	if a == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), a, actual)
	}
}

func TestPerson(t *testing.T) {
	t.Parallel()

	p := PersonStruct{Name: "tsuchinaga", Age: 28}
	actual := Person(p)

	p.Age++
	if p == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), p, actual)
	}
}

func TestPersonAddress(t *testing.T) {
	t.Parallel()

	p := &PersonStruct{Name: "tsuchinaga", Age: 28}
	actual := PersonAddress(p)

	p.Age++
	if p == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), p, actual)
	}
}

func TestPersonStruct_Copy(t *testing.T) {
	t.Parallel()

	p := PersonStruct{Name: "tsuchinaga", Age: 28}
	actual := *p.Copy()

	p.Age++
	if p == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), p, actual)
	}
}

func TestPersonStruct_Copy2(t *testing.T) {
	t.Parallel()

	p := &PersonStruct{Name: "tsuchinaga", Age: 28}
	actual := p.Copy()

	p.Age++
	if p == actual {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), p, actual)
	}
}

func TestSlice(t *testing.T) {
	t.Parallel()

	ns := []int{3}
	actual := Slice(ns)

	ns = append(ns, 6)
	if reflect.DeepEqual(ns, actual) {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), ns, actual)
	}

	if &ns[0] == &actual[0] {
		t.Fatalf("%s 失敗\n元: %v, %p\n先: %v, %p\n", t.Name(), ns[0], &ns[0], actual[0], &actual[0])
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	ns := map[int]int{0: 3}
	actual := Map(ns)

	ns[1] = 6
	if reflect.DeepEqual(ns, actual) {
		t.Fatalf("%s 失敗\n元: %v\n先: %v\n", t.Name(), ns, actual)
	}
}
