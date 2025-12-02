// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import "fmt"

// Add a user named type and provide a comment.
// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare a variable of type user and init using a struct literal.
	u1 := user{
		name:  "Pieter",
		email: "pieter@test.com",
		age:   44,
	}

	// Display the field values.
	fmt.Println("Name:", u1.name)
	fmt.Println("Email:", u1.email)
	fmt.Println("Age:", u1.age)
	fmt.Println()

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "Mabel",
		email: "mabel@test.com",
		age:   4,
	}

	// Display the field values.
	fmt.Println("Name:", u2.name)
	fmt.Println("Email:", u2.email)
	fmt.Println("Age:", u2.age)
}
