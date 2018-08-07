package main

import (
	"fmt"
)

func main() {
	var len, delta int
	var str string
	fmt.Scanf("%d", &len)
	fmt.Scanf("%s", &str)
	fmt.Scanf("%d", &delta)
	var ans []rune
	for _, ch := range str {
		r := ch
		if ch >= 'a' && ch <= 'z' {
			r = rotate(ch, 'a', delta)
		}
		if ch >= 'A' && ch <= 'Z' {
			r = rotate(ch, 'A', delta)
		}
		ans = append(ans, r)
	}
	fmt.Print(string(ans))
}

func rotate(ch rune, base, delta int) rune {
	tmp := int(ch) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}
