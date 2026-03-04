package main

import (
	"fmt"
	"time"
)

// 1. Phương pháp sử dụng vòng lặp For (Iterative)
func fibonacciFor(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var a, b uint64 = 0, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

/* 2. Phương pháp sử dụng đệ quy (Recursive)
func fibonacciRecursive(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
*/
// 3. Phương pháp sử dụng Memorization + Sclice
func fibonacciMemoSclice(n int, memo []uint64) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fibonacciMemoSclice(n-1, memo) + fibonacciMemoSclice(n-2, memo)
	return memo[n]
}
func main() {
	var n int
	fmt.Print("Nhập số nguyên dương n: ")
	_, err := fmt.Scan(&n)

	if err != nil || n < 0 {
		fmt.Println("Vui lòng nhập một số nguyên dương hợp lệ!")
		return
	}

	// 1. Đo thời gian cho phương pháp Vòng lặp For
	startFor := time.Now()
	resultFor := fibonacciFor(n)
	durationFor := time.Since(startFor)
	fmt.Printf("[Vòng lặp For] Kết quả: %d\n", resultFor)
	fmt.Printf("[Vòng lặp For] Thời gian chạy: %v\n", durationFor)

	/* 2. Đo thời gian cho phương pháp Đệ quy
	startRec := time.Now()
	resultRec := fibonacciRecursive(n)
	durationRec := time.Since(startRec)
	fmt.Printf("[Đệ quy]       Kết quả: %d\n", resultRec)
	fmt.Printf("[Đệ quy]       Thời gian chạy: %v\n", durationRec)
	*/
	memo := make([]uint64, n+1)
	// 3. Đo thời gian cho phương pháp Memorization + Sclice
	startMemo := time.Now()
	resultMemo := fibonacciMemoSclice(n, memo)
	durationMemo := time.Since(startMemo)
	fmt.Printf("[Memorization + Sclice]  Kết quả: %d\n", resultMemo)
	fmt.Printf("[Memorization + Sclice]  Thời gian chạy: %v\n", durationMemo)
}
