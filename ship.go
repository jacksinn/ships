package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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

func (ship *Ship) attack(defender *Ship, weapon Weapon) {
	// Generating random damage up to the weapon's max damage
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedDamage := r1.Intn(int(weapon.MaxDamage))
	actualAttackDamage := uint(generatedDamage)

	// Printing out a battle report
	// Letting th euser know who is attacking whom
	fmt.Print("\n", ship.Name, " attacks ", defender.Name, "'s ", ship.Class, " with it's <<", weapon.Name, ">> dealing ", actualAttackDamage, " damage!\n")

	// fmt.Println(ship.Name, "'s", ship.Class, "attacks with it's <<", weapon.Name, ">> dealing", actualAttackDamage, "damage")
	fmt.Println("******Pew pew >>>>>>")

	// Deal the damage to the other player
	// If the attack does more damage than the defender's ship has, then it should die
	if actualAttackDamage >= defender.CurrentHP {
		defender.CurrentHP = 0
	} else {
		defender.takeDamage(actualAttackDamage)
	}

	// Print damage dealt
	defender.printHealth()

	if defender.CurrentHP <= 0 {
		fmt.Println(defender.Name, "'s", defender.Class, "got frigate wrecked!")
		fmt.Println(ship.Name, "is the winner!!!")
		fmt.Println(`
________                        ________                     
/  _____/_____    _____   ____   \_____  \___  __ ___________ 
/   \  ___\__  \  /     \_/ __ \   /   |   \  \/ // __ \_  __ \
\    \_\  \/ __ \|  Y Y  \  ___/  /    |    \   /\  ___/|  | \/
\______  (____  /__|_|  /\___  > \_______  /\_/  \___  >__|   
		\/     \/      \/     \/          \/          \/       
		   					   `)
		os.Exit(0)
	}
	fmt.Println()
}

func (ship *Ship) takeDamage(damage uint) {
	ship.CurrentHP = ship.CurrentHP - damage
}

func (ship Ship) printHealth() {
	// fmt.Println(defender.Name, "has", defender.CurrentHP, "health left")

	fmt.Print("Health: ", ship.CurrentHP, "/", ship.MaxHP, "[", strings.Repeat("*", int(ship.CurrentHP)), strings.Repeat(" ", int(ship.MaxHP-ship.CurrentHP)), "]\n")
}
