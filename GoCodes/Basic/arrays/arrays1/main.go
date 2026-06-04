package main

import "fmt"

func main() {

	name := "Vinay"

	result := reverseString(name)
	fmt.Println(result)
}

func reverseString(name string) string {
	runes := []rune(name)
	length := len(runes)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		// runes[i] = runes[j]
		// runes[j] = runes[i]
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// =================================================

// func reverseString(name string) string {
//     runes := []rune(name)
//     length := len(runes)

//     for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
//         runes[i], runes[j] = runes[j], runes[i]
//     }

//     return string(runes)
// }
