package main

import "fmt"

type Item struct {
	X int
	Y int
}

const (
	maxX = 600
	maxY = 400
)

func main() {
	var i Item
	// It's always good to use the %#v Printf verb for debugging, since it shows types..
	fmt.Printf("i: %#v\n", i)

	// You must specify all the fields in the struct if assigning values only.
	i = Item{10, 20}
	fmt.Printf("i: %#v\n", i)

	// If you are using the field names to assign values to a struct, you can assign them in any order and also partially.
	i = Item{
		Y: 22,
		// X: 11,
	}
	fmt.Printf("i: %#v\n\n", i)

	/* Types of "new" or factory functions.
	func NewItem(x, y int) Item
	func NewItem(x, y int) *Item
	func NewItem(x, y int) (Item, error)
	func NewItem(x, y int) (*Item, error)

	Value semantics: everyone has their own copy.
	Pointer semantics: everyone shares the same copy (on the heap, and might require locking).

	IMPORTANT NOTE: Try to use value semantics wherever possible...
	In general, value semantics is more performant and less memory intensive than pointer semantics (no heap allocation required). */
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, 2000))
}

func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		// Value semantics.
		// return Item{}, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)

		// Pointer semantics.
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}

	i := Item{
		X: x,
		Y: y,
	}

	/* The Go compiler does escape analysis and will allocate i on the heap.
	Run "go run -gcflags=-m game.go" (or "go build -gcflags=-m") in the terminal to see a confirmation of this.
	Again, try to use value semantics wherever possible since value semantics is more performant and less memory intensive than pointer semantics. */
	return &i, nil
}
