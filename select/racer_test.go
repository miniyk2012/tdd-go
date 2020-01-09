package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
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

func TestRacer(t *testing.T) {
	slowServer := makeTestServer(20 * time.Millisecond)
	fastServer := makeTestServer(0 * time.Millisecond)

	slowURL := slowServer.URL // http://127.0.0.1:52446
	fastURL := fastServer.URL // http://127.0.0.1:52447

	want := fastURL
	got := Racer(slowURL, fastURL)

	if want != got {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func makeTestServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return server
}
