package collections

import "fmt"

/*
slices have: pointer, length and capacity.
*/

func Arr_run() {
	arr := []float64{1.99, 2.32, -2.67}
	for idx, num := range arr {
		fmt.Printf("The current number at index %d is %.2f\n", idx, num)
	}
	// a slice can also slice an array
	parentArray := []int{1, 2, 3, 4}
	ss := parentArray[1:]

	// editing the slice also affects the parent array
	ss[0] = 99
	fmt.Printf("ss after editing: %v\n", ss)
	fmt.Printf("parent array after editing: %v\n", parentArray)

	// but if we make the slice point to a different location in memory
	// then the slices pointer will no longer point to the parent arrays location in memory
	ss = append(ss, 901)
	fmt.Printf("ss after adding extra element: %v\n", ss)
	fmt.Printf("parent array after adding extra element to ss: %v\n", parentArray)
}
