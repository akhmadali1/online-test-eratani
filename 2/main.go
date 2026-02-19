package main

import (
	"fmt"
)

// Basic string palindrome check
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test case input
	testInput := []string{"kodokkodok", "jinggaaku", "123321"}

	for _, v := range testInput {
		checkPalindrome := isPalindrome(v)
		outputPalindrome := "itu palindrome"

		if !checkPalindrome {
			outputPalindrome = "itu bukan palindrome"
		}
		fmt.Printf("%s %s\n", v, outputPalindrome)
	}

	// var t int
	// fmt.Scan(&t)

	// k := make([]string, t)
	// for i := 0; i < t; i++ {
	// 	fmt.Scan(&k[i])
	// }

	// for i := 0; i < t; i++ {
	// 	checkPalindrome := isPalindrome(k[i])
	// 	outputPalindrome := "itu palindrome"

	// 	if !checkPalindrome {
	// 		outputPalindrome = "itu bukan palindrome"
	// 	}
	// 	fmt.Printf("%s %s\n", k[i], outputPalindrome)
	// }
}
