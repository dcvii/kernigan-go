package main

import (
	"fmt"
	"unicode/utf8"
)

const GoUsage = `Go is a tool

written in Go.

Usage:
	go command [arguments]

`

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func main() {
	//fmt.Println("Hello, World!")

	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[11])

	//c := s[len(s)-1]
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println("goodbye cruel " + s[7:])

	fmt.Println("marching orders")

	{
		s := "left foot"
		t := s
		s += ", right foot"
		fmt.Println(s)
		fmt.Println(t)
	}

	fmt.Println(GoUsage)

	// now substrings
	fmt.Println("Substring business")
	fmt.Println(HasPrefix("hello", "he"))
	fmt.Println(HasSuffix("hello", "no"))

	s1 := "Hello \x93\xe1" + "Donovan, Alan A. A.; Kernighan, Brian W.. Go Programming Language,(Function). Kindle Edition.  "
	fmt.Println(s1)
	fmt.Println(utf8.RuneCountInString(s1))

	// katakana
	// あいうえお
	s2 := "あいうえお"
	fmt.Printf("%s\n", s2)
}
