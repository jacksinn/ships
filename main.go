package main

import (
	"fmt"
)

// Ship stores information about a ship
type Ship struct {
	Name        string
	Description string
}

func main() {
	fmt.Println(Ship{"Ship 1", "Ship 1 Description"})
}
