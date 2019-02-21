package main

import "time"

func GetMonth(i int) string {
	return time.Month(i).String()
}

func Equal(a, b time.Time) bool {
	return a == b
}

func Gt(a, b time.Time) bool {
	return a.After(b)
}

func Ge(a, b time.Time) bool {
	return !a.Before(b)
}

func Lt(a, b time.Time) bool {
	return !Ge(a, b)
}

func Le(a, b time.Time) bool {
	return !Gt(a, b)
}

func Between(a, b, c time.Time) bool {
	return Ge(a, b) && Le(a, c)
}
