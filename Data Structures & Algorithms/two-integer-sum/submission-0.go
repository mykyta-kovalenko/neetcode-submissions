func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentNumber := nums[i]

		if index, ok := seen[target - currentNumber]; ok {
			return []int{index, i}
		}

		seen[currentNumber] = i
	}

	return []int{}
}
