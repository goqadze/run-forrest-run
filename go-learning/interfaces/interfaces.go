package interfaces

import (
	"fmt"
	"math"
)

// Shape interface demonstrates polymorphism
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Rectangle implements Shape
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// PrintShapeInfo works with any type that implements Shape
func PrintShapeInfo(s Shape) {
	fmt.Printf("  Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

// Writer interface for demonstration
type Writer interface {
	Write(data string) error
}

// ConsoleWriter implements Writer
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) error {
	fmt.Printf("  Writing to console: %s\n", data)
	return nil
}

// FileWriter implements Writer
type FileWriter struct {
	Filename string
}

func (fw FileWriter) Write(data string) error {
	fmt.Printf("  Writing to file '%s': %s\n", fw.Filename, data)
	return nil
}

// RunInterfaces demonstrates interface usage
func RunInterfaces() {
	demonstratePolymorphism()
	demonstrateInterfaceTypes()
	demonstrateEmptyInterface()
}

func demonstratePolymorphism() {
	fmt.Println("Polymorphism with Interfaces:")

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	fmt.Println("  Circle:")
	PrintShapeInfo(circle)

	fmt.Println("  Rectangle:")
	PrintShapeInfo(rectangle)

	// Slice of interfaces
	shapes := []Shape{circle, rectangle}
	fmt.Println("  All shapes:")
	for i, shape := range shapes {
		fmt.Printf("    Shape %d: ", i+1)
		PrintShapeInfo(shape)
	}
}

func demonstrateInterfaceTypes() {
	fmt.Println("Multiple Interface Implementations:")

	writers := []Writer{
		ConsoleWriter{},
		FileWriter{Filename: "output.txt"},
	}

	for _, writer := range writers {
		writer.Write("Hello, Go!")
	}
}

func demonstrateEmptyInterface() {
	fmt.Println("Empty Interface (interface{}):")

	// Empty interface can hold any type
	var anything interface{}

	anything = 42
	fmt.Printf("  Int: %v (type: %T)\n", anything, anything)

	anything = "Hello"
	fmt.Printf("  String: %v (type: %T)\n", anything, anything)

	anything = Circle{Radius: 3}
	fmt.Printf("  Circle: %v (type: %T)\n", anything, anything)

	// Type assertion
	if circle, ok := anything.(Circle); ok {
		fmt.Printf("  Type assertion successful! Radius: %.1f\n", circle.Radius)
	}
}
