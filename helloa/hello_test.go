package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("当输入为空时, 默认提供world", func(tt *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(tt, got, want)
	})

	t.Run("在西班牙", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("在法国", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func TestHello2(t *testing.T) {
	got := Hello("Chris", "")
	want := "Hello, Chris"
	if got != want {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hello("world", "")
	}
}
