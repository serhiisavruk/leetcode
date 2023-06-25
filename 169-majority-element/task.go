package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			m[num] = m[num] + 1
		} else {
			m[num] = 1
		}
	}

	largest := 0
	count := 0

	for key, value := range m {
		if value > count {
			largest = key
			count = value
		}
	}

	return largest
}
