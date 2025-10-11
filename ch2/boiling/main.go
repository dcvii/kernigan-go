package main

import "fmt"

const boilingF, freezingF = 212.0, 32.0

func FToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	fmt.Println("Hello, World!")
	var f = boilingF
	var c = FToC(f)
	fmt.Printf("Freezing = %g F or %g C\n", f, c)

	medals := []string{
		"Gold",
		"Silver",
		"Bronze",
	}

	for i, medal := range medals {
		fmt.Printf("Medal %d: %s\n", i+1, medal)
	}

}
