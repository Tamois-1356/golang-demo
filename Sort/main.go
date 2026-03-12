package main

import (
	"fmt"
)

func inputSlice() []int {
	var size int
	fmt.Print("Nhập số lượng phần tử của mảng: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Printf("Nhập phần tử thứ %d: ", i+1)
		fmt.Scan(&arr[i])
	}
	return arr
}

func main() {

	// Bubble Sort
	data1 := inputSlice()

	if len(data1) > 0 {
		fmt.Printf("Mảng vừa nhập: %v\n", data1)
		BubbleSort(data1)
		fmt.Printf("Kết quả Bubble Sort: %v\n", data1)
	}

	// Quick Sort
	data2 := inputSlice()

	if len(data2) > 0 {
		fmt.Printf("Mảng vừa nhập: %v\n", data2)
		QuickSort(data2, 0, len(data2)-1)
		fmt.Printf("Kết quả Quick Sort: %v\n", data2)
	}
}
