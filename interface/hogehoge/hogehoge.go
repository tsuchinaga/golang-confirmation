package hogehoge

type HogeHogeInterface interface {
	Hoge() string
}

type HogeHoge struct {
	HogeHogeInterface
}

func (h *HogeHoge) Hoge() string {
	return "HOGEHGOE"
}
