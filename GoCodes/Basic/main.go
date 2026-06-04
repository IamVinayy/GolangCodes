package main

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

func main() {
	var num1 int8
	var num2 uint8

	num1 = math.MaxInt8
	num2 = math.MaxUint8

	fmt.Println(num1, num2)
	color.Red("Prints text in Color.")
	color.White("Prints text in Color.")
	color.Green("Prints text in Color.")

}
