package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Ship stores information about a ship
type Ship struct {
	ID          uint
	Name        string
	Description string
	MaxWeapons  uint
	Weapons     []Weapon
	Class       string // Cruiser, Caravel, Frigate, Dreadnaught
	MaxHP       uint
	CurrentHP   uint
	Color       string
	MaxSpeed    uint

	// Ship may have shields and other attachments
}

func (ship *Ship) attack(defender *Ship, weapon Weapon) {
	// Letting th euser know who is attacking whom
	fmt.Println(ship.Name, " is attacking ", defender.Name)

	// Generating random damage up to the weapon's max damage
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedDamage := r1.Intn(int(weapon.MaxDamage))
	actualAttackDamage := uint(generatedDamage)

	fmt.Println(ship.Name, "'s", ship.Class, "attacks with it's <<", weapon.Name, ">> dealing", actualAttackDamage, "damage")
	fmt.Println("******Pew pew >>>>>>")
	// Deal the damage to the other player
	// If the attack does more damage than the defender's ship has, then it should die
	if actualAttackDamage >= defender.CurrentHP {
		defender.CurrentHP = 0
	} else {
		defender.takeDamage(actualAttackDamage)
	}

	fmt.Println(defender.Name, "has", defender.CurrentHP, "left")
	if defender.CurrentHP <= 0 {
		fmt.Println(defender.Name, "'s", defender.Class, "got frigate wrecked!")
	}
	fmt.Println("--------------------------")
}

func (ship *Ship) takeDamage(damage uint) {
	ship.CurrentHP = ship.CurrentHP - damage
}

// Weapon stores information about a weapon
type Weapon struct {
	ID             uint
	Name           string
	Description    string
	Cost           uint
	MaxDamage      uint
	Distance       uint
	MaxShots       uint
	RemainingShots uint
}

// Player information
type Player struct {
	ID    uint
	Name  string
	Level uint
}

func main() {
	// Player 1
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
	}

	// Player 2
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
	}

	// over 9000
	over9k := Weapon{
		ID:             9000,
		Name:           "over 9000",
		Description:    "This weapon was made by the saltyest of people.",
		Cost:           9000,
		MaxDamage:      75,
		Distance:       100,
		MaxShots:       1,
		RemainingShots: 1,
	}

	// widow maker
	widowMaker := Weapon{
		ID:             1000,
		Name:           "Widow Maker",
		Description:    "Time to bury yo spouse.",
		Cost:           9000,
		MaxDamage:      75,
		Distance:       100,
		MaxShots:       1,
		RemainingShots: 1,
	}

	// ion accelerator
	ionAccelerator := Weapon{
		ID:             9001,
		Name:           "Ion Accelerator",
		Description:    "Uh it accelerates ions. Duh.",
		Cost:           100,
		MaxDamage:      75,
		Distance:       100,
		MaxShots:       1,
		RemainingShots: 1,
	}
	// corn dog
	cornDog := Weapon{
		ID:             1001,
		Name:           "Corn Dog",
		Description:    "A tasty treat for your defeat.",
		Cost:           100,
		MaxDamage:      75,
		Distance:       100,
		MaxShots:       1,
		RemainingShots: 1,
	}

	// Attaching weapons to ships
	dad.Weapons = []Weapon{widowMaker, cornDog}
	steven.Weapons = []Weapon{over9k, ionAccelerator}

	for i := 1; i <= 5; i++ {
		fmt.Println("===========================")
		fmt.Println("Round", i, "::FIGHT!!!!!!!!")
		// Dad Turn 1
		dad.attack(&steven, dad.Weapons[0])

		// Steven turn 1
		steven.attack(&dad, steven.Weapons[0])
	}

}
