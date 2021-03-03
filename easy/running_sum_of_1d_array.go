package easy

func RunningSum(numbers []int) []int {
	var result []int
	currentNumber := 0
	for _, v := range numbers {
		currentNumber += v
		result = append(result, currentNumber)
	}
	return result
}
