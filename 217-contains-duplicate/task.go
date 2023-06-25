package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			return true
		}
		m[num] = num
	}
	return false
}
