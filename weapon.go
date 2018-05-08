package main

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
