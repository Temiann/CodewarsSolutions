package main

import (
	"strings"
	"unicode"
)

func CamelCase(s string) string {
	if len(s) == 0 {
		return s
	}
	var new_str string
	arr := strings.Split(s, " ")
	for _, value := range arr {
		for i, value2 := range value {
			if i == 0 {
				new_str += string(unicode.ToTitle(value2))
			} else {
				new_str += string(value2)
			}
		}
	}
	return new_str
}
