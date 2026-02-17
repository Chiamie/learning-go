package main

import "fmt"

func main() {

	fmt.Print("Enter account number: ")
	var accountNumber int
	fmt.Scan(&accountNumber)
	
	fmt.Print("What is your balance at the beginning of the month: ")
	var beginningBalance int
	fmt.Scan(&beginningBalance)
	
	fmt.Print("Enter charges for the month: ")
	var charges int
	fmt.Scan(&charges)
	
	fmt.Print("Enter the total credits applied for this month: ")
	var credit int
	fmt.Scan(&credit)
	
	const creditLimit int = 5000
	

	
	output := beginningBalance + charges - credit
	fmt.Printf("The new balance is %d\n", output)
	if output > creditLimit{
		fmt.Print("Credit limit exceeded")
	}
	
	
	
	
	
	
}
