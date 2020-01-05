package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("测试调用次数", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}

		Countdown(buffer, spySleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
		if spySleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
		}
	})

	t.Run("sleep after every print", func(t *testing.T) {
		spy := &CountdownOperationsSpy{}
		Countdown(spy, spy)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spy.Calls) {
			t.Errorf("wanted calls %v got %v", want, spy.Calls)
		}
	})
}
