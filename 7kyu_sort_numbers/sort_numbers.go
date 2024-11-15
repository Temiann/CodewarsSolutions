package main

func SortNumbers(numbers []int) []int {
	if len(numbers) > 0 {
		for i := 0; i < len(numbers)-1; i++ {
			for j := 0; j < len(numbers)-i-1; j++ {
				if numbers[j+1] < numbers[j] {
					numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				}
			}
		}
		return numbers
	} else {
		return []int{}
	}
}
