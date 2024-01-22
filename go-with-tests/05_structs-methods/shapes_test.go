package main

import (
    "testing"
)

func TestPerimeter(t *testing.T) {
    assertCorrectPerimeter := func(t testing.TB, got, want float64) {
        t.Helper()
        if got != want {
            t.Errorf("expected %v but got %v", want, got)
        }
    }
    t.Run("Rectangle Perimeter", func(t *testing.T) {
        rectangle := Rectangle{10.0, 10.0}
        got := rectangle.Perimeter()
        want := 40.0

        assertCorrectPerimeter(t, got, want)
    })
}

func TestArea(t *testing.T) {
//    checkArea := func(t testing.TB, shape Shape, want float64) {
//        t.Helper()
//        got := shape.Area()
//        assertCorrectArea(t, got, want)
//    }
    areaTests := []struct {
        name string
        shape Shape
        want float64
    }{
        // Named fields
        // "The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations.
        {name: "Rectangle", shape: Rectangle{width: 10.0, length: 10.0}, want: 100},
        {name: "Circle", shape: Circle{radius: 10}, want: 314.1592653589793},
        {name: "Triangle", shape: Triangle{base: 12, height: 6}, want: 37.0},
    }

    for _, tt := range areaTests {
        t.Run(tt.name, func(t *testing.T) {
            got := tt.shape.Area()
            if got != tt.want {
                t.Errorf("%#v expected %v but got %v", tt.shape, tt.want, got)
            }
        })
    }
}
