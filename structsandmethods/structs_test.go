package structsandmethods

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{7, 12}
	got := Perimeter(rectangle)
	want := 38.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		t.Helper()

// 		got := shape.Area()

// 		if got != want {
// 			t.Errorf("Got %g want %g", got, want)
// 		}
// 	}

// 	t.Run("Rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{7, 12}
// 		checkArea(t, rectangle, 84.0)
// 	})

// 	t.Run("Circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v Got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
