package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("world")
	}
}
