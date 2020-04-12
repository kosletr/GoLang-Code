package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	const N = 10000
	const max = 10000

	a := make([]int, N)

	initArr(&a, N, max)

	fmt.Println(a)

	qsort(&a, 0, len(a)-1)

	fmt.Println(a)

}

func qsort(arr *[]int, low int, high int) {

	if low < high {
		pivot := partition(arr, low, high)

		go qsort(arr, low, pivot-1)
		go qsort(arr, pivot+1, high)
		time.Sleep(1)
	}

}

func partition(arr *[]int, low int, high int) int {

	i := low
	var j int = low

	for j = low; j < high; j++ {

		if (*arr)[j] < (*arr)[high] {
			swap(arr, i, j)
			i++
		}

	}
	swap(arr, i, j)
	return i
}

func swap(arr *[]int, i int, j int) {

	temp := (*arr)[i]
	(*arr)[i] = (*arr)[j]
	(*arr)[j] = temp

}

func initArr(arr *[]int, N int, max int) {
	for i := 0; i < N; i++ {
		(*arr)[i] = rand.Intn(max)
	}
}
