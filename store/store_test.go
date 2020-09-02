package store

import (
	"fmt"
	"testing"
	"time"
)

func Test_1(t *testing.T) { // 結果は常に一致する
	a := GetFoo()
	b := GetFoo()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	a.Open()
	b.Close()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	a.Close()
	b.Open()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = true, b = true

	a.Open()
	b.Open()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = true, b = true

	a.Close()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false
}

func Test_2(t *testing.T) { // 非同期でも結果は常に一致する
	a := GetFoo()
	b := GetFoo()
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	go a.Open()
	go b.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	go a.Close()
	go b.Open()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	go a.Open()
	go b.Open()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = true, b = true

	go a.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false
}

func Test_3(t *testing.T) { // それぞれ違う待機時間があってもfoo.mtxによって一致させられる
	a := GetFoo().(*foo)
	b := GetFoo().(*foo)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	a.isWait = true
	a.wait = 50 * time.Millisecond

	b.isWait = true
	b.wait = 75 * time.Millisecond

	go a.Open()
	go b.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	go a.Close()
	go b.Open()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false

	go a.Open()
	go b.Open()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = true, b = true

	go a.Close()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("a = %v, b = %v\n", a.IsOpen(), b.IsOpen()) // a = false, b = false
}
