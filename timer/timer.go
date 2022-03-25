package timer

import (
	"fmt"
	"github.com/mark2185/pomogoro/util"
)

const (
	TOMATO = `🍅`
	BREAK  = `🏖`
	PAUSE  = "\u23f8"
)

type State int8

const (
	Break State = iota
	Work
)

type Timestamp struct {
	minutes int
	seconds int
}

func GetTimestampFromSeconds(seconds int) Timestamp {
	minutes, seconds := util.SecondsToMinutes(seconds)
	return Timestamp{minutes, seconds}
}

func (t *Timestamp) UpdateTime(deltaMinutes, deltaSeconds int) {
	t.minutes += deltaMinutes
	t.seconds += deltaSeconds
	t.minutes, t.seconds = util.SecondsToMinutes(t.GetTimeInSeconds())
}

func (t *Timer) UpdateTime(deltaMinutes, deltaSeconds int) {
	t.timestamp.UpdateTime(deltaMinutes, deltaSeconds)
}

func (t *Timestamp) GetTimeInSeconds() int {
	return t.seconds + t.minutes*60
}

func (t *Timer) Tick() {
	t.timestamp.UpdateTime(0, -1)
}

func (t *Timer) ToString() string {
	icon := ""
	switch t.GetState() {
	case Work:
		icon = TOMATO
	case Break:
		icon = BREAK
	}
	return fmt.Sprintf("%s %02d:%02d", icon, t.timestamp.minutes, t.timestamp.seconds)
}

type Timer struct {
	timestamp Timestamp
	state     State
	running   bool
}

func (t *Timer) IsRunning() bool {
	return t.running
}

func (t *Timer) GetSeconds() int {
	return t.timestamp.seconds
}

func (t *Timer) GetMinutes() int {
	return t.timestamp.minutes
}

func (t *Timer) Resume() {
	t.running = true
}

func (t *Timer) Pause() {
	t.running = false
}

func (t *Timer) Toggle() {
	t.running = !t.running
}

func (t *Timer) Stop() {
	t.Reset()
	t.Pause()
}

func (t *Timer) Reset() {
	switch t.state {
	case Work:
		t.timestamp = Timestamp{25, 0}
	case Break:
		t.timestamp = Timestamp{5, 0}
	}
}

func (t *Timer) GetState() State {
	return t.state
}

func (t *Timer) Switch() {
	switch t.state {
	case Work:
		t.state = Break
	case Break:
		t.state = Work
	}
	t.Reset()
}

func MakeTimer() Timer {
	return Timer{Timestamp{25, 0}, Work, false}
}
