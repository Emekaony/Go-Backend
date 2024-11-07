package collections

import "fmt"

// just know that in go, slices mean arrays that vary in length
// this is not the same as slices in, say Rust for example which is a reference to an underlying array
// or vector

func Arr_run() {
	arr := []float64{1.99, 2.32, -2.67}
	for idx, num := range arr {
		fmt.Printf("The current number at index %d is %.2f\n", idx, num)
	}
}
