package clock_interface

type Timer struct {
	clock ClockInterface
}

func (t *Timer) Message() string {
	switch t.clock.Now().Hour() {
	case 9, 10, 11, 13, 14, 15, 16, 17:
		return "働け"
	case 12:
		return "昼飯くうでー"
	case 18:
		return "帰んでー"
	default:
		return "暇やでー"
	}
}
