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
			t.Errorf("%#v got %.2f want %.2f",shape, got, want)
		}
	}

	t.Run("rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72)
	})

	t.Run("circle Area", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("表格驱动测试", func(t *testing.T) {
		// 创建匿名结构体切片
		areaTest := []struct {
			shape Shape
			want  float64
		}{
			{shape: Rectangle{Width: 12, Height: 6}, want: 72},
			{Circle{10}, 314.1592653589793},
			{Triangle{12, 6}, 35.0},
		}
		for _, tt := range areaTest {
			checkArea(t, tt.shape, tt.want)
		}
	})
}
