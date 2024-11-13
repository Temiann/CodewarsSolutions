package main

import (
	"strings"
)

// solution
func wave(words string) []string {
	var arr []string
	for i, value := range words {
		if value != ' ' {
			res := ""
			for x, value2 := range words {
				if i == x {
					res = res + strings.ToUpper(string(value2))
				} else {
					res = res + string(value2)
				}
			}
			arr = append(arr, res)
		}
	}
	return arr
}
