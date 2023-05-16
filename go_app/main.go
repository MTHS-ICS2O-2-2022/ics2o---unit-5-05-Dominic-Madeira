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
	var lengthA float64
	var lengthB float64
	var lengthC float64

	// input
	fmt.Println("This program finds what type of triangle you have.")
	fmt.Println()
	fmt.Print("Enter your first side length: ")
	fmt.Scanln(&lengthA)
	fmt.Println()
	fmt.Print("Enter your second side length: ")
	fmt.Scanln(&lengthB)
	fmt.Println()
	fmt.Print("Enter your third side length: ")
	fmt.Scanln(&lengthC)
	fmt.Println()

	// process
	// cosine law
	angleA := math.Acos((math.Pow(lengthB, 2)+math.Pow(lengthC, 2)-math.Pow(lengthA, 2))/(2*lengthB*lengthC)) * (180 / math.Pi)
	angleB := math.Acos((math.Pow(lengthC, 2)+math.Pow(lengthA, 2)-math.Pow(lengthB, 2))/(2*lengthC*lengthA)) * (180 / math.Pi)
	angleC := math.Acos((math.Pow(lengthA, 2)+math.Pow(lengthB, 2)-math.Pow(lengthC, 2))/(2*lengthA*lengthB)) * (180 / math.Pi)

	sumOfAngles := angleA + angleB + angleC
	sumOfAngles = math.Round(sumOfAngles)

	angleA = math.Round(angleA)
	angleB = math.Round(angleB)
	angleC = math.Round(angleC)

	if sumOfAngles == 180 {
		if angleA != angleB && angleA != angleC && angleB != angleC {
			if angleA == 90 || angleB == 90 || angleC == 90 {
				fmt.Println("You have a right angle triangle")
			} else {
				fmt.Println("You have a scalene triangle")
			}
		} else {
			if angleA == angleB && angleA == angleC {
				fmt.Println("You have an equilateral triangle")
			} else {
				if angleA == 90 || angleB == 90 || angleC == 90 {
					fmt.Println("You have a right angle triangle")
				} else {
					fmt.Println("You have an isosceles triangle")
				}
			}
		}
	} else {
		fmt.Println("Please enter a valid triangle")
	}

	fmt.Println()
	fmt.Println("\nDone.")
}
