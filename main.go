package main

import "fmt"

func main() {
	// 2D array that represents the rows and columns of the airplane
	seats := [][]int{{3, 4}, {4, 5}, {2, 3}, {3, 4}}

	// Number of passengers waiting in queue
	numPassengers := 30

	// Variable to keep track of the number of seats filled
	seatCounter := 0

	// Loop through the 2D array
	for row := 0; row < len(seats); row++ {
		for col := 0; col < seats[row][0]; col++ {
			// Check if seat is an aisle seat
			if col == 0 || col == seats[row][0]-1 {
				seats[row][col] = 1
				seatCounter++
			} else if col == 1 {
				// Check if seat is a window seat
				seats[row][col] = 1
				seatCounter++
			} else if seatCounter < numPassengers {
				// Fill the center seat
				seats[row][col] = 1
				seatCounter++
			}
			// check if all seats have been filled
			if seatCounter == numPassengers {
				break
			}
		}
	}

	// Print out the final seating arrangement
	fmt.Println(seats)
}
