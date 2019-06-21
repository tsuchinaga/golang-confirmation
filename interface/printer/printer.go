package printer

import (
	"gitlab.com/tsuchinaga/golang-confirmation/interface/hogehoge"
	"log"
)

type HogeInterface interface {
	Hoge() string
}

func HogePrinter(hoge HogeInterface) {
	log.Println(hoge.Hoge())
}

func HogeHogePrinter(hoge hogehoge.HogeHogeInterface) {
	log.Println(hoge.Hoge())
}
