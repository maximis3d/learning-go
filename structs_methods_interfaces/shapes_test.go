package structsmethodsinterfaces

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 4}, hasArea: 24.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {

	areaTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasPerimeter: 36},
		{name: "Circle", shape: Circle{Radius: 10}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, Side1: 10, Side2: 10}, hasPerimeter: 32},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}
