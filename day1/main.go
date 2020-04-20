package main

import (
	"fmt"
	"strconv"

	helpers "../helpers"
)

func main() {
	puzzleInput := helpers.GetFile("./input.txt")
	var intialFuel int
	var totalFuelRequired int

	for _, item := range puzzleInput {
		input, err := strconv.ParseFloat(item, 64)
		if err != nil {
			panic(err)
		}

		calcultedFuelPerInput := getRequiredFuelByMass(input)
		intialFuel += calcultedFuelPerInput
		totalFuelRequired += getRecursiveFuel(input)

	}

	fmt.Printf("Intial fuel required is %v \n and total required fule is %v \n", intialFuel, totalFuelRequired)
}

func getRequiredFuelByMass(input float64) int {
	requiredFuel := int((input / 3) - 2)
	return requiredFuel
}

func getRecursiveFuel(fuelInput float64) int {
	var totalExtraFuel = getRequiredFuelByMass(fuelInput)
	var remainingMass = totalExtraFuel

	for remainingMass > 0 {
		calculateExtraFuel := getRequiredFuelByMass(float64(remainingMass))

		if calculateExtraFuel <= 0 {
			break
		}

		totalExtraFuel += calculateExtraFuel
		remainingMass = calculateExtraFuel
	}

	return totalExtraFuel
}
