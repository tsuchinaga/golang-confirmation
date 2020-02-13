package copy

// String - stringを複製する
func String(s string) string {
	return s
}

// Int - intを複製する
func Int(n int) int {
	return n
}

// Float - floatを複製する
func Float(f float64) float64 {
	return f
}

// AgeStruct - int型のstruct
type AgeStruct int

// Age - AgeStructを複製する
func Age(a AgeStruct) AgeStruct {
	return a
}

// AgeStruct_Copy - AgeStructの複製を作る
func (a *AgeStruct) Copy() AgeStruct {
	return *a
}

// PersonStruct - 名前と年齢を持つ人の構造体
type PersonStruct struct {
	Name string
	Age  int
}

// Person - PersonStructを複製する
func Person(p PersonStruct) PersonStruct {
	return p
}

// PersonAddress - PersonStructのアドレスを複製する
func PersonAddress(p *PersonStruct) *PersonStruct {
	// return p すると同じアドレスで返される

	p2 := *p
	return &p2
}

// PersonStruct_Copy - PersonStructの複製を作る
func (p *PersonStruct) Copy() *PersonStruct {
	// return p すると同じアドレスで返される

	p2 := *p
	return &p2
}

// Slice - sliceを複製する
func Slice(ns []int) []int {
	nss := make([]int, len(ns))
	copy(nss, ns)
	return nss
}

// Map - mapを複製する
func Map(ns map[int]int) map[int]int {
	// return ns すると同じアドレスで返される

	ns2 := make(map[int]int)
	for k, v := range ns {
		ns2[k] = v
	}

	return ns2
}
