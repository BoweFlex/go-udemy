package main

import "math"

type Rectangle struct {
    width float64
    length float64
}

func (rectangle Rectangle) Perimeter() (perimeter float64) {
    perimeter = 2 * (rectangle.width + rectangle.length)
    return
}

func (rectangle Rectangle) Area() (area float64) {
    area = rectangle.length * rectangle.width
    return
}

type Circle struct {
    radius float64
}

func (circle Circle) Area() (area float64) {
    area = math.Pi * circle.radius * circle.radius
    return
}

type Triangle struct {
    base float64
    height float64
}

func (triangle Triangle) Area() (area float64) {
    return (triangle.base * triangle.height) * 0.5
}

type Shape interface {
    Area() float64
}

