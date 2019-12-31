package day01

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Calculate fuel requirements for a mass
func calculateFuelForMass(mass int) int {
	return int(math.Floor(float64(mass)/3.0)) - 2
}

// Calculate fuel requirement for a module (incl. fuel for fuel)
func calculateFuelForModule(moduleMass int) int {
	overallFuelForModule := 0

	for fuel := moduleMass; fuel > 0; {
		fuel = calculateFuelForMass(fuel)

		if fuel > 0 {
			overallFuelForModule += fuel
		}
	}

	return overallFuelForModule
}

// CalculateFuel calculates the fuel requirements for the entire ship
func CalculateFuel() {
	file, err := os.Open("../../resources/day01_masses.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	overallFuelStar1, overallFuelStar2 := 0, 0
	for scanner.Scan() {
		massForModule, _ := strconv.Atoi(scanner.Text())

		overallFuelStar1 += calculateFuelForMass(massForModule)
		overallFuelStar2 += calculateFuelForModule(massForModule)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Required fuel (Star 1): %v\n", overallFuelStar1)
	fmt.Printf("Required fuel (Star 2): %v\n", overallFuelStar2)
}
