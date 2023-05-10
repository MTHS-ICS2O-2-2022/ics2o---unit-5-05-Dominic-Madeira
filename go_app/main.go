// Created by: Dominic M.
// Created on: April 2023
//
// This program finds what type of triangle you have.
package main

import (
	"fmt"
	"math"
)

func main() {
	var sideA int
	var sideB int
	var sideC int

	// input
	fmt.Println("This program finds what type of triangle you have.")
	fmt.Println()
	fmt.Print("Enter your first side length: ")
	fmt.Scanln(&sideA)
	fmt.Println()
	fmt.Print("Enter your second side length: ")
	fmt.Scanln(&sideB)
	fmt.Println()
	fmt.Print("Enter your third side length: ")
	fmt.Scanln(&sideC)

	// process
	// cosine law
	angleA := (sideB * sideB) + (sideC * sideC) - (sideA * sideA) / (2 * sideB * sideC) * (180 / math.Pi)
	angleB := (sideA * sideA) + (sideC * sideC) - (sideB * sideB) / (2 * sideA * sideC) * (180 / math.Pi)
	angleC := (sideA * sideA) + (sideB * sideB) - (sideC * sideC) / (2 * sideA * sideB) * (180 / math.Pi)






	fmt.Println()
	fmt.Println("\nDone.")
}
