package basics

import "fmt"

/*
challenge 1:
	Write a filter function that takes a slice of int and a predicate.
	The filter function returns a slice with values in the original slice
	for which the predicate returns true.
*/

func Run() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 15, 20, 35, 22, 16, 90}
	oddSlice := filter(isOdd, slice)
	evenSlice := filter(isEven, slice)
	divisibleByFiveSlice := filter(dividibleByFive, slice)
	fmt.Printf("Odd slice: %v.\n", oddSlice)
	fmt.Printf("Even slice: %v.\n", evenSlice)
	fmt.Printf("Divisible by five slice: %v.\n", divisibleByFiveSlice)
}

func filter(predicate func(int) bool, values []int) []int {
	// pretty simple function to be honest
	var newSlice []int
	for _, item := range values {
		if predicate(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func isOdd(num int) bool {
	return num%2 != 0
}

func dividibleByFive(num int) bool {
	return num%5 == 0
}

func isEven(num int) bool {
	return num%2 == 0
}
