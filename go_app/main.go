// Created by: Bonnie Z
// Created on: April 2023
//
// This program finds the area and perimeter of a rectangle

package main

import "fmt"

func main() {
	// This function finds the volume of a right rectangular prism
	var pay int
	var hours int
	var tax int
	var payment int

	// input
	fmt.Println("This program to find the payment and taxes")
	fmt.Println()
	fmt.Print("Enter the hours: ")
	fmt.Scanln(&hours)
	fmt.Print("Enter the pay: ")
	fmt.Scanln(&pay)
	fmt.Print("Enter the tax: ")
	fmt.Scanln(&tax)

	// process
	payment = (hours * pay) * (1.00 - 0.18)

	// output
	fmt.Println()
	fmt.Println("`Your pay will be: $${payment.toFixed(2)}`")
	fmt.Println("\nDone.")
}
