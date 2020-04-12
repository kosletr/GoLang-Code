package main

import (
	//"fmt"
	"fmt"
	"math/rand"
)

func main() {

	rand.Seed(12)

	const N = 4

	var a [N]float64

	initArr(a[:], 15)

	printArr(a[:])

	qSort(a[:])

	printArr(a[:])

}

func initArr(arr []float64, max int) []float64 {
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Float64() * float64(max)
	}
	return arr
}

func printArr(arr []float64) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%.2f ", arr[i])
	}
	fmt.Print("\n\n")
}

func qSort(arr []float64) []float64 {

	swap := func(arr []float64, i int, j int) {

		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp

	}

	for i := 1; i < len(arr); i++ {
		for j := len(arr) - 1; j >= i; j-- {
			if arr[j-1] > arr[j] {
				swap(arr[:], j-1, j)
			}
			fmt.Println("Rep: ", i)
			printArr(arr[:])
		}
	}
	return arr
}
