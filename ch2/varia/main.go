package main

import (
	"fmt"
	"log"
	"os"
)

const boilingF, freezingF = 212.0, 32.0

func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	fmt.Println("Hello, World!")
	var f = blah()
	var c = FToC(f)
	fmt.Printf("Freezing = %v F or %g C\n", f, c)

	var fspec, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer fspec.Close()
	fmt.Println("File opened successfully.")

	if derf() {
		fmt.Println("derf returned true")
	}
}

func blah() float64 {
	var i = boilingF
	return i
}

func derf() bool {
	var x, y int
	fmt.Println(&x == &y)
	log.Println("derf called")
	return (&x == &y)
}
