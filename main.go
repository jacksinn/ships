package main

import (
	"fmt"
)

// Ship stores information about a ship
type Ship struct {
	Name        string
	Description string
	MaxWeapons  uint
	Class       string // Cruiser, Frigate, etc
	MaxHP       uint
}

// Weapon stores information about a weapon
type Weapon struct {
	Name        string
	Description string
	Cost        uint
	Power       uint
	Damage      uint
	Range       uint
}

func main() {
	fmt.Println(Ship{"Ship 1", "Ship 1 Description", 2, "Frigate", 100})
}
