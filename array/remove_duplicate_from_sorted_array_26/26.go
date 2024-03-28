package remove_duplicate_from_sorted_array_26

func removeDuplicates(nums []int) int {
	lastUniIndex := 0
	for i, num := range nums {
		if i == 0 || nums[i-1] != nums[i] {
			nums[lastUniIndex] = num
			lastUniIndex++
		}
	}
	return lastUniIndex
}
