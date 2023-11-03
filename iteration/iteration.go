package main

import "fmt"

func main() {
	fmt.Println(Repeat("", 4))
}

func Repeat(char string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += char
	}
	return repeated
}
