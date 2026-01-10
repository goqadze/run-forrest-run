package basics

import "fmt"

// RunBasics demonstrates fundamental Go concepts
func RunBasics() {
	// Variables and types
	demonstrateVariables()

	// Control structures
	demonstrateControlFlow()

	// Functions
	demonstrateFunction()

	// Slices and maps
	demonstrateCollections()
}

func demonstrateVariables() {
	fmt.Println("Variables and Types:")

	// Different ways to declare variables
	var name string = "Go"
	age := 15 // Short declaration
	var isAwesome bool = true
	const pi = 3.14159

	fmt.Printf("  Language: %s, Age: %d, Awesome: %t, Pi: %.2f\n", name, age, isAwesome, pi)
}

func demonstrateControlFlow() {
	fmt.Println("Control Flow:")

	// If-else
	x := 10
	if x > 5 {
		fmt.Println("  x is greater than 5")
	}

	// For loop (only loop in Go!)
	fmt.Print("  Counting: ")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Switch
	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("  Start of the work week")
	case "Friday":
		fmt.Println("  Almost weekend!")
	default:
		fmt.Println("  Another day")
	}
}

func demonstrateFunction() {
	fmt.Println("Functions:")

	// Multiple return values
	sum, product := calculate(4, 5)
	fmt.Printf("  Sum: %d, Product: %d\n", sum, product)

	// Named return values
	result := namedReturn(10)
	fmt.Printf("  Named return result: %d\n", result)
}

// Multiple return values
func calculate(a, b int) (int, int) {
	return a + b, a * b
}

// Named return values
func namedReturn(x int) (result int) {
	result = x * 2
	return // Naked return
}

func demonstrateCollections() {
	fmt.Println("Collections:")

	// Slices (dynamic arrays)
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	fmt.Printf("  Slice: %v\n", numbers)

	// Maps (key-value pairs)
	person := map[string]string{
		"name": "Alice",
		"city": "NYC",
	}
	fmt.Printf("  Map: %v\n", person)
	fmt.Printf("  Name from map: %s\n", person["name"])
}
