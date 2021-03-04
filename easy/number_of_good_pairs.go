package easy

func NumIdenticalPairs(nums []int) int {
	counter := 0
	indexA := 0
	for indexA < len(nums) {
		indexB := indexA + 1
		for indexB < len(nums) {
			if nums[indexA] == nums[indexB] {
				counter++
			}
			indexB++
		}
		indexA++
	}
	return counter
}
