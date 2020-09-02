package store

import (
	"sync"
	"time"
)

var (
	fooStore Foo
	fooMtx   sync.Mutex
)

func GetFoo() Foo {
	fooMtx.Lock()
	defer fooMtx.Unlock()

	if fooStore == nil {
		fooStore = &foo{}
	}
	return fooStore
}

type Foo interface {
	IsOpen() bool
	Open()
	Close()
}

type foo struct {
	isOpen bool
	isWait bool
	wait   time.Duration
	mtx    sync.Mutex
}

func (f *foo) IsOpen() bool {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	return f.isOpen
}

func (f *foo) Open() {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if f.isWait {
		time.Sleep(f.wait)
	}
	f.isOpen = true
}

func (f *foo) Close() {
	f.mtx.Lock()
	defer f.mtx.Unlock()
	if f.isWait {
		time.Sleep(f.wait)
	}
	f.isOpen = false
}
