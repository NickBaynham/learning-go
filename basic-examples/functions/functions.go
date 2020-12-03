package main

import (
	"fmt"
	"log"
)

// Global Scope
var metersPerLiter float64

func main() {
	metersPerLiter = 10.0
	var width, height, area float64
	
	sayHi()
	repeatLine("Hello!", 3)

	// first wall
	width = 4.2
	height = 3.0
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/metersPerLiter)

	// second wall
	width = 5.2
	height = 3.5
	area = width * height
	fmt.Printf("%.2f liters needed\n", area/metersPerLiter)

	// replaced duplicate code with function calls

	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)

	fmt.Printf("%.2f liters needed\n", getPaintNeeded(5.5, 2.9))

	// with error handling

	amount, err := getPaintNeededWithErrorHandling(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%0.2f liters needed\n", amount)
	}

	// using pointers

	value := 6
	double(&value)
	fmt.Println(value)
}

// define a new function

func sayHi() {
	fmt.Println("Hi!")
}

// define a function with arguments

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}

// replacing the repetitive code with a function

func paintNeeded(width float64, height float64) {
	// local scope
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/metersPerLiter)
}

// function with return values

func getPaintNeeded(width float64, height float64) float64 {
	area := width * height
	return area / metersPerLiter
}

// function with an error return vaue

func getPaintNeededWithErrorHandling(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / metersPerLiter, nil
}

// function with pointers

func double (number *int) {
	*number *= 2
}