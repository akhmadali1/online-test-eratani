package main

import "fmt"

func main() {
	// Precompute valid numbers (1 <= k <= 1000)
	valid := make([]int, 0, 1000)
	for i := 1; len(valid) < 1000; i++ {
		if i%3 != 0 && i%10 != 3 {
			valid = append(valid, i)
		}
	}

	// Test case input
	// testInput := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000}
	
	// for _, v := range testInput {
	// 	fmt.Println(valid[v-1])
	// }

	var t int
	fmt.Scan(&t)

	k := make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Scan(&k[i])
	}

	for i := 0; i < t; i++ {
		fmt.Println(valid[k[i]-1])
	}
}
