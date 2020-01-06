package racer

import (
	"testing"
)

func TestRace(t *testing.T) {
	fastURL := "http://www.baidu.com"
	slowURL := "http://www.sina.cn"

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
