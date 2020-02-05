package main

import (
	"fmt"
)

func main() {

	xs := []float64{1, 2, 3, 4}
	avg := Average(xs)
	fmt.Println(avg)

}

/*Average takes the input of an array of numbers and returns their average. */
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
