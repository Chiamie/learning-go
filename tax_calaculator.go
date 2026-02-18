package main

import "fmt"

func main() {

	var annualSalary, totalTax float64


	fmt.Print("Enter your annual salary in USD: ")
	fmt.Scan(&annualSalary)
	
	if annualSalary <= 30000{
		totalTax = annualSalary * 0.15
	} else {
		totalTax = annualSalary * 0.20
	}
	
	fmt.Printf("Your total tax is %.2f\n", totalTax)
	
}