package main

import (
	"fmt"
)

func main() {
	var value string
	answere := 1
	fmt.Scanf("%s", &value)
	for _, ch := range value {
		if ch >= 'A' && ch <= 'Z' {
			answere++
		}
	}
	fmt.Print(answere)
}
