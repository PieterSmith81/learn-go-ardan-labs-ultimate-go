// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p player) average() float64 {
	if p.atBats == 0 {
		return 0
	}

	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{"Pieter", 100, 25},
		{"Ida", 50, 10},
		{"Mabel", 75, 25},
		{
			name:   "Willow",
			atBats: 5,
		},
	}

	// Display the batting average for each player in the slice.
	for i := range players {
		fmt.Printf("Player: %s\t\t\tBatting Average: .%.f\n", players[i].name, players[i].average()*1000)
	}
}
