package fuga

type HogeInterface interface {
	Hoge() string
}

type Fuga struct {
	HogeInterface
}

func (h *Fuga) Hoge() string {
	return "FUGA"
}
