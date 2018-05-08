package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Player information
type Player struct {
	ID      uint
	Name    string
	Level   uint
	Balance uint
	Ship    Ship
}

func (attacker *Player) attack(defender *Player, weapon Weapon) {
	// Generating random damage up to the weapon's max damage
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedDamage := r1.Intn(int(weapon.MaxDamage))
	actualAttackDamage := uint(generatedDamage)

	// Printing out a battle report
	// Letting th euser know who is attacking whom
	fmt.Print("\n", attacker.Name, " attacks ", defender.Name, "'s ", defender.Ship.Class, " with it's <<", weapon.Name, ">> dealing ", actualAttackDamage, " damage!\n")

	// fmt.Println(ship.Name, "'s", ship.Class, "attacks with it's <<", weapon.Name, ">> dealing", actualAttackDamage, "damage")
	fmt.Println("******Pew pew >>>>>>")

	// Deal the damage to the other player
	// If the attack does more damage than the defender's ship has, then it should die
	if actualAttackDamage >= defender.Ship.CurrentHP {
		defender.Ship.CurrentHP = 0
	} else {
		defender.Ship.takeDamage(actualAttackDamage)
	}

	// Print damage dealt
	defender.Ship.printHealth()

	if defender.Ship.CurrentHP <= 0 {
		fmt.Println(defender.Name, "'s", defender.Ship.Class, "got frigate wrecked!")
		fmt.Println(attacker.Name, "is the winner!!!")
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
