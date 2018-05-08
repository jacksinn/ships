package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// Player 1
	dad := Player{
		ID:      1,
		Name:    "Jack",
		Level:   1,
		Balance: 1000,
	}
	dad.Ship = Ship{
		ID:          1,
		Name:        "Frigate I",
		Description: "Ship 1 Description",
		MaxWeapons:  2,
		Class:       "Frigate I",
		MaxHP:       100,
		CurrentHP:   100,
		Color:       "Grey",
		MaxSpeed:    100,
	}

	// Player 2
	steven := Player{
		ID:      1,
		Name:    "Steven",
		Level:   1,
		Balance: 1000,
	}

	steven.Ship = Ship{
		ID:          2,
		Name:        "Dreadnaught I",
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
		MaxDamage:      30,
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
		MaxDamage:      30,
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
		MaxDamage:      5,
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
		MaxDamage:      5,
		Distance:       100,
		MaxShots:       1,
		RemainingShots: 1,
	}

	// Attaching weapons to ships
	dad.Ship.Weapons = []Weapon{widowMaker, cornDog}
	steven.Ship.Weapons = []Weapon{over9k, ionAccelerator}

	for i := 1; ; i++ {
		fmt.Println(strings.Repeat("=", 80))
		fmt.Println("Round", i, "::FIGHT!!!!!!!!")

		// Randomize the weapon chosen
		// Generating random damage up to the weapon's max damage
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		dadsWeapon := dad.Ship.Weapons[r1.Intn(len(dad.Ship.Weapons))]

		// Dad Turn
		dad.attack(&steven, dadsWeapon)

		// Steven turn
		stevensWeapon := steven.Ship.Weapons[r1.Intn(len(steven.Ship.Weapons))]
		steven.attack(&dad, stevensWeapon)
		fmt.Println(strings.Repeat("=", 80))
	}

}
