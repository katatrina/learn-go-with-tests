package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Width:  4.0,
		Height: 5.0,
	}
	got := rectangle.Perimeter()
	want := 18.0
	
	if got != want {
		t.Errorf("got %.2f expectedArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expectedArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}
	
	for _, subtest := range areaTests {
		t.Run(subtest.name, func(t *testing.T) {
			received := subtest.shape.Area()
			if received != subtest.expectedArea {
				t.Errorf("%#v received %g expected %g", subtest.shape, received, subtest.expectedArea)
			}
		})
	}
}
