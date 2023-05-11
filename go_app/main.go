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
	fmt.Println()

	// process
	// cosine law
	angleA := (sideB * sideB) + (sideC * sideC) - (sideA * sideA) / (2 * sideB * sideC) * (180 / math.Pi)
	angleB := (sideA * sideA) + (sideC * sideC) - (sideB * sideB) / (2 * sideA * sideC) * (180 / math.Pi)
	angleC := (sideA * sideA) + (sideB * sideB) - (sideC * sideC) / (2 * sideA * sideB) * (180 / math.Pi)

	angleAFormatted := fmt.Sprintf("%.2f", angleA)
	angleBFormatted := fmt.Sprintf("%.2f", angleB)
	angleCFormatted := fmt.Sprintf("%.2f", angleC)

	sumOfAngles := angleAFormatted + angleBFormatted + angleCFormatted

	if sumOfAngles == "180" {
		if angleAFormatted != angleBFormatted && angleAFormatted != angleCFormatted && angleBFormatted != angleCFormatted {
			if angleAFormatted == 90 || angleBFormatted == 90 || angleCFormatted == 90{
				fmt.Println("You have a right angle triangle")
			} else {
				fmt.Println("You have a scalene triangle")
			}
		} else {
			if angleAFormatted == angleBFormatted && angleAFormatted == angleCFormatted {
				fmt.Println("You have an equilateral triangle")
			} else {
				if angleAFormatted == 90 || angleBFormatted == 90 || angleCFormatted == 90{
					fmt.Println("You have a right angle triangle")
				} else {
					fmt.Println("You have an isosceles triangle")
			}
		}
	} else {
		fmt.Println("Please enter a valid triangle")
	}

	fmt.Println()
	fmt.Println("\nDone.")
}
