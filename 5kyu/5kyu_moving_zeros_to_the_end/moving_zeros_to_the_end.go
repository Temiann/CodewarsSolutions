package kata

func MoveZeros(arr []int) []int {
	arrWithoutZero := make([]int, 0, 10)
	arrZero := make([]int, 0, 10)
	for _, value := range arr {
		if value != 0 {
			arrWithoutZero = append(arrWithoutZero, value)
		} else {
			arrZero = append(arrZero, value)
		}
	}
	newArr := append(arrWithoutZero, arrZero...)
	return newArr
}
