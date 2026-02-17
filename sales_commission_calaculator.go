package main

import "fmt"

func main() {
		items := map[int]float64{
		1: 239.99,
		2: 129.75,
		3: 99.95,
		4: 350.89,
	}

	totalSalesValue := 0.0
	continueInput := "yes"

	
	for continueInput == "yes" {
		itemNum, quantity := prompt()
		
		
		price := items[itemNum]
		totalSalesValue += price * float64(quantity)

		fmt.Print("Did you sell another item (yes/no)? ")
		fmt.Scan(&continueInput)
	}

	
	grossSalesCommission := 0.09 * totalSalesValue
	totalEarnings := 200.0 + grossSalesCommission

	fmt.Printf("\nTotal Gross Sales: $%.2f\n", totalSalesValue)
	fmt.Printf("Total Earnings: $%.2f\n", totalEarnings)
}

func prompt() (int, int) {
	var itemSold, itemQuantity int

	for {
		fmt.Print("Enter item number (1-4): ")
		fmt.Scan(&itemSold)
		if itemSold >= 1 && itemSold <= 4 {
			break
		}
		fmt.Println("Invalid Item! Please enter 1, 2, 3, or 4.")
	}

	for {
		fmt.Printf("How many of item %d did you sell? ", itemSold)
		fmt.Scan(&itemQuantity)
		if itemQuantity > 0 {
			break
		}
		fmt.Println("Invalid quantity! Must be greater than 0.")
	}

	return itemSold, itemQuantity
}
