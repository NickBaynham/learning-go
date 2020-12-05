// examples of defining types
package main

import (
	"fmt"
)

type liters float64
type gallons float64

func main() {
	var carFuel gallons
	var busFuel liters
	carFuel = gallons(10.00)
	busFuel = liters(240.00)
	fmt.Println(carFuel, busFuel)
}
