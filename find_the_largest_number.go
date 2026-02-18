/*
Pseudo code:

collect input of 10 integers
determine the largest integer
output the largest integer to the console
*/




package main

import "fmt"


func main(){

	var number int
	var numbers [10] int
	count := 0
	for count < 10{
		fmt.Print("Enter a whole number: ")
		fmt.Scan(&number)
		numbers[count] = number
		count += 1
	}
	
	largest := getTheLargestNumberIn(numbers)
	
	
	
	fmt.Printf("\nThe largest number you entered is %d\n", largest)
	
	

}

func getTheLargestNumberIn(numbers [10]int)int{
	
	
	var largest int = numbers[0]
	for index := 1; index < len(numbers); index++{
		if numbers[index] > largest{
			largest = numbers[index]
		}
	}
	return largest

}





