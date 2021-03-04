package easy

func ShuffleTheArray(nums []int) []int {
	length := len(nums)
	result := make([]int, length)
	counter := 0
	for counter < length/2 {
		result[counter*2] = nums[counter]
		result[counter*2+1] = nums[counter+length/2]
		counter++
	}
	return result
}
