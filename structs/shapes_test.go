package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 20}
	got := rectangle.Perimeter()
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})
	
	t.Run("circle Area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.159265358973)
	})
	
}