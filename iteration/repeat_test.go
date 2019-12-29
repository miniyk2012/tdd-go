package iteration

import (
	"reflect"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func TestStings(t *testing.T) {
	t.Run("test strings.Contains", func(t *testing.T) {
		if !strings.Contains("yangkai", "yangk") {
			t.Errorf("%q not contains %q", "yangkai", "yangk")
		}
	})

	t.Run("test strings.Fields", func(t *testing.T) {
		a := strings.Fields("  foo bar  baz   ")
		b := [] string{"foo", "bar", "baz"}

		if !reflect.DeepEqual(a, b) {
			t.Errorf("Fields not correct")
		}
	})
}
