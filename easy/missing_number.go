package easy

func MissingNumber(numbers []int) int {
	total := 0
	totalAll := 0
	for i, v := range numbers {
		total += v
		totalAll += i + 1
	}
	return totalAll - total
}

