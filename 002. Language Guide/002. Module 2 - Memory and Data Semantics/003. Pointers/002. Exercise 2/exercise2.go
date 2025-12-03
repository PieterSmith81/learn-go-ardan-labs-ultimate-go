// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"

// Declare a type named user.
// user represents a user in the system.
type user struct {
	name  string
	email string
	age   int
}

// Create a function that changes the value of one of the user fields.
// Note that this is an example of using pointer semantics (to mutate the value of one of the fields in a struct).
func modifyAge(user *user, newAge int) { // add pointer parameter, add value parameter.

	// Use the pointer to change the value that the
	// pointer points to.
	user.age = newAge
}

func main() {

	// Create a variable of type user and initialize each field.
	u1 := user{
		name:  "Pieter",
		email: "pieter@test.com",
		age:   44,
	}

	// Display the value of the variable.
	fmt.Println("Name:", u1.name, "\tEmail:", u1.email, "\tAge:", u1.age)

	// Share the variable with the function you declared above.
	modifyAge(&u1, 45)

	// Display the value of the variable.
	fmt.Println("Name:", u1.name, "\tEmail:", u1.email, "\tAge:", u1.age)
}
