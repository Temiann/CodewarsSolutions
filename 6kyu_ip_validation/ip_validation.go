package main

import (
	"strconv"
	"strings"
	"unicode"
)

func Is_valid_ip(ip string) bool {
	if ip == "" || strings.Count(ip, ".") != 3 {
		return false
	}
	arr := strings.Split(ip, ".")
	if len(arr) != 4 {
		return false
	}
	for _, value := range arr {
		if len(value) == 0 || (len(value) > 1 && value[0] == '0') {
			return false
		}
		for _, value2 := range value {
			if !unicode.IsDigit(value2) {
				return false
			}
		}
		num, err := strconv.Atoi(value)
		if err != nil || num < 0 || num > 255 {
			return false
		}
	}
	return true
}
