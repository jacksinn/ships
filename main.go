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
	CurrentHP   uint
	Color       string
	MaxSpeed    uint
	MaxAttack   uint

	// Ship may have shields and other attachments
}

func (ship *Ship) attack(defender *Ship) {
	fmt.Println(ship.Name, " is attacking ", defender.Name)

	fmt.Println(ship.Name, "attacks with his weapon dealing", ship.MaxAttack, "damage")
	fmt.Println("******Pew pew >>>>>>")
	// Deal the damage to the other player
	// If the attack does more damage than the defender's ship has, then it should die
	if ship.MaxAttack >= defender.CurrentHP {
		defender.CurrentHP = 0
	} else {
		defender.takeDamage(ship.MaxAttack)
	}

	fmt.Println(defender.Name, "has", defender.CurrentHP, "left")
	if defender.CurrentHP <= 0 {
		fmt.Println(defender.Name, "'s shipped got frigate wrecked!")
	}
	fmt.Println("--------------------------")
}

func (ship *Ship) takeDamage(damage uint) {
	ship.CurrentHP = ship.CurrentHP - damage
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
		CurrentHP:   100,
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
		CurrentHP:   100,
		MaxAttack:   50,
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("===========================")
		fmt.Println("Round", i, "::FIGHT!!!!!!!!")
		// Dad Turn 1
		dad.attack(&steven)

		// Steven turn 1
		steven.attack(&dad)
	}

}
