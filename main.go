package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
		MaxDamage:      25,
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
		MaxDamage:      25,
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

		// Randomize the weapon chosen
		// Generating random damage up to the weapon's max damage
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		dadsWeapon := dad.Weapons[r1.Intn(len(dad.Weapons))]

		// Dad Turn
		dad.attack(&steven, dadsWeapon)

		// Steven turn
		stevensWeapon := steven.Weapons[r1.Intn(len(steven.Weapons))]
		steven.attack(&dad, stevensWeapon)
	}

}
