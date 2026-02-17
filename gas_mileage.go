package main

import "fmt"

func main() {

	fmt.Print("Enter the miles driven: ")
	var milesDriven int
	fmt.Scan(&milesDriven)
	
	fmt.Print("Enter the number of gallons used: ")
	var numberOfGallons int
	fmt.Scan(&numberOfGallons)
	
	

	
	output := milesDriven / numberOfGallons
	fmt.Printf("The miles per gallon used for this trip is %d\n", output)
	
	
	
	
	
	
}