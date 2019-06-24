package clock_interface

import "time"

// 時計のインターフェース
type ClockInterface interface {
	Now() time.Time
}

// 普段使う用の時計
type Clock struct{}

func (c *Clock) Now() time.Time {
	return time.Now()
}

// テストで使う用の時計
type ClockAtTest struct {
	CurrentTime time.Time
}

func (c *ClockAtTest) Now() time.Time {
	return c.CurrentTime
}
