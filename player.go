package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player information
type Player struct {
	ID          uint
	Name        string
	Level       uint
	Balance     uint
	Ships       []Ship
	CurrentShip Ship
	Wins        uint
	Losses      uint
}

func (attacker *Player) attack(defender *Player, weapon Weapon) bool {
	// Generating random damage up to the weapon's max damage
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	generatedDamage := r1.Intn(int(weapon.MaxDamage))
	actualAttackDamage := uint(generatedDamage)

	// Printing out a battle report
	// Letting th euser know who is attacking whom
	fmt.Print("\n", attacker.Name, " attacks ", defender.Name, "'s ", defender.CurrentShip.Class, " with it's <<", weapon.Name, ">> dealing ", actualAttackDamage, " damage!\n")

	// fmt.Println(ship.Name, "'s", ship.Class, "attacks with it's <<", weapon.Name, ">> dealing", actualAttackDamage, "damage")
	fmt.Println("******Pew pew >>>>>>")

	// Deal the damage to the other player
	// If the attack does more damage than the defender's ship has, then it should die
	if actualAttackDamage >= defender.CurrentShip.CurrentHP {
		defender.CurrentShip.CurrentHP = 0
	} else {
		defender.CurrentShip.takeDamage(actualAttackDamage)
	}

	// Print damage dealt
	defender.CurrentShip.printHealth()

	if defender.CurrentShip.CurrentHP <= 0 {
		fmt.Println(defender.Name, "'s", defender.CurrentShip.Class, "got frigate wrecked!")
		fmt.Println(attacker.Name, "is the winner!!!")
		fmt.Println(`
________                        ________                     
/  _____/_____    _____   ____   \_____  \___  __ ___________ 
/   \  ___\__  \  /     \_/ __ \   /   |   \  \/ // __ \_  __ \
\    \_\  \/ __ \|  Y Y  \  ___/  /    |    \   /\  ___/|  | \/
\______  (____  /__|_|  /\___  > \_______  /\_/  \___  >__|   
		\/     \/      \/     \/          \/          \/       
								  `)
		return true
	}
	fmt.Println()
	return false
}
