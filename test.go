package main

import "fmt"

func main() {

	arr := make([]int, 3)
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	algo(&arr)
	fmt.Printf("%v ", arr)
}

func algo(arr *[]int) {
	(*arr)[1] = 10
	*arr = append(*arr, 15)
}
