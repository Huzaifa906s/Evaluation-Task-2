package main

import (
	"fmt"
	"sort"
)

func waveSort(arr []int, x int) {
	n := len(arr)
	blockSize := 2*x + 1

	if n%blockSize != 0 {
		fmt.Println("Error: Array length must be a multiple of 2x+1")
		return
	}

	for i := 0; i < n; i += blockSize {

		block := arr[i : i+blockSize]

		sort.Ints(block)

		maxIdx := len(block) - 1
		block[x], block[maxIdx] = block[maxIdx], block[x]

	}
}

func main() {
	arr := []int{3, 6, 5, 10, 7, 1, 8, 4, 9, 2, 11, 12, 15, 14, 13}
	x := 2

	fmt.Println("Original Array:", arr)
	waveSort(arr, x)
	fmt.Println("Rearranged Array:", arr)
}
