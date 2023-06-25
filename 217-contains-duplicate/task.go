package main

func containsDuplicate(nums []int) bool {
	hasZero := false
	m := make(map[int]int)
	for _, num := range nums {
		if num == 0 {
			if hasZero {
				return true
			}
			hasZero = true
			continue
		}

		if m[num] == num {
			return true
		}
		m[num] = num
	}
	return false
}
