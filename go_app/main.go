// createby:Bonnie ZHu
// Date: 2023 april
// Description: This is a simple program to calculate the pay and tax

package main

import (
	"fmt"
)

const taxRate = 0.18

func main() {
	// input
	var hours, rate int
	fmt.Print("Enter the number of hours worked: ")
	fmt.Scanln(&hours)
	fmt.Print("Enter the hourly rate: ")
	fmt.Scanln(&rate)
	fmt.Println() // add an empty line here

	// proccess
	pay := float64(hours*rate) * (1 - taxRate)
	tax := float64(hours*rate) * taxRate

	// output
	fmt.Printf("Your pay will be: $%.2f\n", pay)
	fmt.Printf("The government will take: $%.2f\n", tax)

	fmt.Println("\nDone.")
}
