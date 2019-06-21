package hoge

type HogeInterface interface {
	Hoge() string
}
type Hoge struct{}

func (h *Hoge) Hoge() string {
	return "HOGE"
}
