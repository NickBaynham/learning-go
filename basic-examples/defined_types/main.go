// examples of defining types
package main

import (
	"fmt"
	"github.com/nickbaynham/fuelTypes"
)

func main() {
	carFuel := fuelTypes.Gallons(10.00)
	busFuel := fuelTypes.Liters(240.00)
	fmt.Println(carFuel, busFuel)
	myFuel := carFuel.ToLiters(carFuel)
	fmt.Println(myFuel)

	// Using the FuelTypes defined types
	car1 := fuelTypes.Gallons(10.00)
	car2 := fuelTypes.Liters(240.00)
	fmt.Println(car1.ToLiters(car1), car2.ToGallons(car2))
}
