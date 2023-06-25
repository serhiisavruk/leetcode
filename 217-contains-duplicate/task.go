package main

func containsDuplicate(nums []int) bool {
	dict := make([]int, 0)
	for _, num := range nums {
		for _, d := range dict {
			if d == num {
				return true
			}
		}
		dict = append(dict, num)
	}
	return false
}
