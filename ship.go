package main

import (
	"fmt"
	"strings"
)

/*
 *  0 - Corvette
 *  1 - Freighter
 *  2 - Frigate
 *  3 - Cruiser
 *  4 - Carrier
 *  5 - Dreadnought
 */

// Ship stores information about a ship
type Ship struct {
	ID          uint
	Name        string
	Description string
	MaxWeapons  uint
	Weapons     []Weapon
	Class       string
	MaxHP       uint
	CurrentHP   uint
	Color       string
	MaxSpeed    uint

	// Ship may have shields and other attachments
}

func (ship *Ship) takeDamage(damage uint) {
	ship.CurrentHP = ship.CurrentHP - damage
}

func (ship Ship) printHealth() {
	// fmt.Println(defender.Name, "has", defender.CurrentHP, "health left")

	fmt.Print("Health: ", ship.CurrentHP, "/", ship.MaxHP, "[", strings.Repeat("*", int(ship.CurrentHP)), strings.Repeat(" ", int(ship.MaxHP-ship.CurrentHP)), "]\n")
}
