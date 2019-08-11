package base

import "time"

type TimeFactory struct{}
type Timer time.Timer
type Ticker time.Ticker

func (TimeFactory) Now() time.Time {
	return time.Now()
}

func Duration(ms int) time.Duration {
	return time.Duration(ms) * time.Millisecond
}

func (TimeFactory) Unix() int64 {
	return time.Now().Unix()
}

func (TimeFactory) Zero() time.Time {
	return time.Time{}
}

func (TimeFactory) UnixNano() int64 {
	return time.Now().UnixNano()
}

func (TimeFactory) Time(args ...int) time.Time {
	var val = [4]int{}
	if len(args) > 4 {
		args = args[:4]
	}
	for i, v := range args {
		val[i] = v
	}
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, val[0], val[1], val[2], val[3], time.Local)
}

func (TimeFactory) Date(args ...int) time.Time {
	var val = [7]int{0, 1, 1}
	if len(args) > 7 {
		args = args[:7]
	}
	for i, v := range args {
		val[i] = v
	}
	return time.Date(val[0], time.Month(val[1]), val[2], val[3], val[4], val[5], val[6], time.Local)
}

func (TimeFactory) Format(layout string) string {
	return time.Now().Format(layout)
}

func (TimeFactory) Parse(layout, value string) (time.Time, error) {
	return time.Parse(layout, value)
}

func (f TimeFactory) ParseUnix(layout, value string) int64 {
	tm, _ := f.Parse(layout, value)
	return tm.Unix()
}

func (f TimeFactory) ParseUnixNano(layout, value string) int64 {
	tm, _ := f.Parse(layout, value)
	return tm.UnixNano()
}

func (TimeFactory) After(s int, callback func()) *Timer {
	return (*Timer)(time.AfterFunc(Duration(s), callback))
}

func (TimeFactory) Sleep(s int) {
	time.Sleep(Duration(s))
}

func (TimeFactory) Timer(s int) *Timer {
	return (*Timer)(time.NewTimer(Duration(s)))
}

func (TimeFactory) Ticker(s int) *Ticker {
	return (*Ticker)(time.NewTicker(Duration(s)))
}

func (t *Timer) Wait() bool {
	_, ok := <-t.C
	return ok
}
func (t *Timer) Close() {
	t.Stop()
}
func (t *Timer) Stop() bool {
	return (*time.Timer)(t).Stop()
}
func (t *Timer) Reset(s int) bool {
	return (*time.Timer)(t).Reset(Duration(s))
}

func (t *Ticker) Wait() bool {
	_, ok := <-t.C
	return ok
}
func (t *Ticker) Close() {
	t.Stop()
}
func (t *Ticker) Stop() {
	(*time.Ticker)(t).Stop()
}
