package main

import (
	"fmt"
	"slices"
)

type Item struct {
	X int
	Y int
}

const (
	maxX = 600
	maxY = 400
)

type Key byte

const (
	Copper Key = iota + 1
	Jade
	Crystal
)

type Player struct {
	Name string
	Item // Player embeds Item.
	Keys []Key
}

func main() {
	// Struct creation and initialization examples.
	var i Item
	// It's always good to use the %#v Printf verb for debugging, since it shows types.
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
	fmt.Println()

	// Pointer vs value semantics (and mutating a function parameter).
	i.Move(10, 20)
	fmt.Printf("i after move: %#v\n\n", i)

	// Embedding example.
	p1 := Player{
		Name: "Parzival",
	}
	fmt.Printf("p1: %+v\n", p1)
	fmt.Println("p1.X:", p1.X) // or "fmt.Println("p1.Item.X:", p1.Item.X)" if you need more specificity.
	p1.Move(100, 200)
	fmt.Printf("p1 after move: %+v\n\n", p1)

	/* Exercise:
	- Add a "Keys" field to Player which is a slice of strings.
	- Add a "Found(key string)" method to player.
		- It should err if key is not one of "jade", "copper", or "crystal".
		- It should add a key only once. */
	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Copper)) // <nil>
	fmt.Println(p1.Found(Key(7))) // unknown key
	fmt.Println("keys:", p1.Keys) // keys: [copper]
	fmt.Println()

	// Interfaces
	ms := []Mover{
		&i,
		&p1,
	}

	moveAll(ms, 50, 70)
	for _, m := range ms {
		fmt.Println(m)
	}
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

/*
Move moves i by delta x and delta y.

"i" is called the "receiver".
i is a pointer receiver.
Value vs pointer receiver
- In general use value semantics.
- Try to keep the same semantics on all methods.
- But there are some cases when you must use a pointer receiver:
  - If you have a lock field.
  - If you need to mutate the struct (like in the function below).
  - Decoding/unmarshalling.
*/
func (i *Item) Move(dx, dy int) {
	i.X += dx
	i.Y += dy
}

// Found verifies if a valid key was found for a player or not.
func (p *Player) Found(key Key) error {
	switch key {
	case Copper, Jade, Crystal:
		// OK
	default:
		return fmt.Errorf("unknown key %q", key)
	}

	if !slices.Contains(p.Keys, key) {
		p.Keys = append(p.Keys, key)
	}

	return nil
}

/*
	Interfaces

- Set of methods (and types).
- We define interfaces as "what you need", not "what you provide".
  - Interfaces in Go are small (the Go standard library's average is around 2 methods per interface).
  - If you have an interface with more than 4 methods, think again (i.e., split it into smaller interfaces).

- Best practice: Start with concrete types, then discover the interfaces.

- Rule of thumb: Accept interfaces, return types.
*/
type Mover interface {
	Move(int, int)
}

// moveAll moves the x and y deltas for one or more (i.e., multiple) Item/Player structs.
func moveAll(ms []Mover, dx, dy int) {
	for _, m := range ms {
		m.Move(dx, dy)
	}
}

/*
	Interfaces thought exercise: Sorting

	func Sort(s Sortable) {
		// ...
	}

	type Sortable interface {
		Less(i, j int) bool
		Swap(i, j int)
		Len() int
	}
*/

// String implements the fmt.Stringer interface.
func (k Key) String() string {
	switch k {
	case Copper:
		return "copper"
	case Jade:
		return "jade"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}
