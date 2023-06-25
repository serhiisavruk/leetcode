package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			return true
		}
		m[num] = true
	}
	return false
}
