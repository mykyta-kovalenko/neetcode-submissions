func topKFrequent(nums []int, k int) []int {
	frequencyMap := make(map[int]int)
	for _, num := range nums {
		frequencyMap[num]++
	}

	buckets := make([][]int, len(nums) + 1)
	for num, count := range frequencyMap {
		buckets[count] = append(buckets[count], num)
	}

	result := make([]int, 0, k)
	for i := len(nums); i >= 1; i-- {
		for _, num := range buckets[i] {
			result = append(result, num)
			if len(result) == k {
				return result
			}
		}
			
	}

	return []int{}
}
