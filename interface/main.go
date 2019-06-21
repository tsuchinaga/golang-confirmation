package main

import (
	"gitlab.com/tsuchinaga/golang-confirmation/interface/fuga"
	"gitlab.com/tsuchinaga/golang-confirmation/interface/hoge"
	"gitlab.com/tsuchinaga/golang-confirmation/interface/hogehoge"
	"gitlab.com/tsuchinaga/golang-confirmation/interface/printer"
	"log"
)

type HogeInterface interface {
	Hoge() string
}

type Hoge struct{}

func (h *Hoge) Hoge() string {
	return "HOGE"
}

type Fuga struct {
	HogeInterface
}

func (h *Fuga) Hoge() string {
	return "FUGA"
}

type Piyo struct {
	HogeInterface
}

func main() {
	log.Println(new(Hoge).Hoge())
	log.Println(new(Fuga).Hoge())
	// log.HogePrinter(new(Piyo).Hoge()) // panic: runtime error: invalid memory address or nil pointer dereference

	log.Println("-----")

	printer.HogePrinter(new(hoge.Hoge))
	printer.HogePrinter(new(Hoge))
	printer.HogePrinter(new(fuga.Fuga))
	printer.HogePrinter(new(Fuga))
	// printer.HogePrinter(new(Piyo)) // panic: runtime error: invalid memory address or nil pointer dereference

	log.Println("-----")
	printer.HogeHogePrinter(new(hogehoge.HogeHoge))
	printer.HogeHogePrinter(new(hoge.Hoge))
	printer.HogeHogePrinter(new(Hoge))
	printer.HogeHogePrinter(new(fuga.Fuga))
	printer.HogeHogePrinter(new(Fuga))
}
