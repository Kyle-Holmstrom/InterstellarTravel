package main

import "fmt"

// Create the function fuelGauge()
func fuelGauge(fuel int) {
	fmt.Println("Amount of fuel left:", fuel)
}

// Create the function calculateFuel()
func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	default:
		fuel = 0
	}
	return fuel
}

// Create the function greetPlanet()
func greetPlanet(planet string) {
	fmt.Println("Welcome to ", planet)
}

// Create the function cantFly()
func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet()
func flyToPlanet(planet string, fuel int) int {
	var fuelRemaining, fuelCost int
	fuelRemaining = fuel
	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else if fuelCost > fuelRemaining {
		cantFly()
	}
	return fuelRemaining
}

func main() {
	// Create `planetChoice` and `fuel`
	var fuel int
	fuel = 1000000
	planetChoice := "Venus"
	// And then liftoff!
	fuel = flyToPlanet(planetChoice, fuel)
	fuelGauge(fuel)
}
