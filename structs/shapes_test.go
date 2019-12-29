package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 20}
	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	
	t.Run("circle Area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.16

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	
}