package main

import (
	"fmt"
)

// Ship stores information about a ship
type Ship struct {
	ID          uint
	Name        string
	Description string
	MaxWeapons  uint
	Class       string // Cruiser, Caravel, Frigate, Dreadnaught
	MaxHP       uint
	Color       string
	MaxSpeed    uint
	MaxAttack   uint

	// Ship may have shields and other attachments
}

func (attacker Ship) attack(defender Ship) {
	fmt.Println(attacker.Name, " is attacking ", defender.Name)
}

// Weapon stores information about a weapon
type Weapon struct {
	ID          uint
	Name        string
	Description string
	Cost        uint
	Power       uint
	Damage      uint
	Range       uint
}

// Player information
type Player struct {
	ID    uint
	Name  string
	Level uint
}

func main() {
	dad := Ship{
		ID:          1,
		Name:        "Dad",
		Description: "Ship 1 Description",
		MaxWeapons:  2,
		Class:       "Frigate",
		MaxHP:       100,
		Color:       "Grey",
		MaxSpeed:    100,
		MaxAttack:   50,
	}

	steven := Ship{
		ID:          2,
		Name:        "Steven",
		Description: "Ship 2 Description",
		MaxWeapons:  2,
		Class:       "Dreadnaught",
		MaxHP:       100,
		Color:       "Red",
		MaxSpeed:    100,
		MaxAttack:   50,
	}

	dad.attack(steven)

}
