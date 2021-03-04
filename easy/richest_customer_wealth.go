package easy

func RichestCustomerWealth(account [][]int) int {
	result := 0
	for _, customer := range account {
		total := 0
		for _, v := range customer {
			total += v
		}
		if total > result {
			result = total
		}
	}
	return result
}
