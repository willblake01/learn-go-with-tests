package shapes

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}

	got := Perimeter(rectangle.Width, rectangle.Height)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		tt.shape.Area()

		if tt.shape.Area() != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, tt.shape.Area(), tt.want)
		}
	}
}
