package easy

// LargestAltitude finding the largest altitude
func LargestAltitude(gain []int) int {
	largest := 0
	current := 0
	for _, v := range gain {
		current += v
		if current > largest {
			largest = current
		}
	}
	return largest
}
