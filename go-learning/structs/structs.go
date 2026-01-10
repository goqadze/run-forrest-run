package structs

import "fmt"

// Person demonstrates a basic struct
type Person struct {
	Name string
	Age  int
	City string
}

// Employee demonstrates struct embedding
type Employee struct {
	Person   // Embedded struct
	JobTitle string
	Salary   float64
}

// Method on Person struct (value receiver)
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s", p.Name, p.City)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) HaveBirthday() {
	p.Age++
}

// RunStructs demonstrates structs and methods
func RunStructs() {
	// Creating structs
	demonstrateBasicStruct()

	// Struct methods
	demonstrateMethods()

	// Struct embedding
	demonstrateEmbedding()

	// Pointers
	demonstratePointers()
}

func demonstrateBasicStruct() {
	fmt.Println("Basic Structs:")

	// Different ways to create structs
	person1 := Person{Name: "Alice", Age: 30, City: "NYC"}
	person2 := Person{"Bob", 25, "LA"} // Order matters
	var person3 Person                 // Zero values
	person3.Name = "Charlie"
	person3.Age = 35
	person3.City = "Chicago"

	fmt.Printf("  %+v\n", person1)
	fmt.Printf("  %+v\n", person2)
	fmt.Printf("  %+v\n", person3)
}

func demonstrateMethods() {
	fmt.Println("Struct Methods:")

	person := Person{Name: "David", Age: 28, City: "Boston"}
	fmt.Printf("  %s\n", person.Greet())

	// Pointer receiver method
	person.HaveBirthday()
	fmt.Printf("  After birthday: Age = %d\n", person.Age)
}

func demonstrateEmbedding() {
	fmt.Println("Struct Embedding:")

	emp := Employee{
		Person:   Person{Name: "Eve", Age: 32, City: "Seattle"},
		JobTitle: "Software Engineer",
		Salary:   95000.50,
	}

	// Can access embedded fields directly
	fmt.Printf("  Employee: %s, Job: %s\n", emp.Name, emp.JobTitle)
	// Can also call embedded methods
	fmt.Printf("  %s\n", emp.Greet())
}

func demonstratePointers() {
	fmt.Println("Pointers:")

	person := Person{Name: "Frank", Age: 40, City: "Austin"}
	fmt.Printf("  Before: %+v\n", person)

	// Pass pointer to modify original
	modifyPerson(&person)
	fmt.Printf("  After:  %+v\n", person)
}

func modifyPerson(p *Person) {
	p.Age = 41
	p.City = "Dallas"
}
