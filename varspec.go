package main

import "fmt"

var slice = []float64{1, 2, 3, 4, 5}
var total float64
var converter int

func findAverage(x []float64) []float64 {
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	converter := float64(len(slice))
	fmt.Println(total / converter)
	fmt.Printf("%T \n", converter)
	return x
}

func main() {
	fmt.Printf("%T \n", converter)
	findAverage(slice)
	fmt.Printf("%T \n", converter)
}
