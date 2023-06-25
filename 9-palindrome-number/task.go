package main

import (
	"math"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	len := lenLoop10(x)
	power := float64(len - 1)
	start := int(math.Pow(10, power))
	end := 10
	start2 := 1
	for i := 0; i < len/2; i++ {
		a1 := x / start
		a2 := x % end
		b1 := getLast(a1)
		b2 := a2 / start2
		end = end * 10
		start = start / 10
		start2 = start2 * 10
		//print(b1, b2)
		if b1 != b2 {
			return false
		}
	}
	return true
}

func getLast(x int) int {
	if x < 10 {
		return x
	}
	x = x % 10
	return getLast(x)
}

func lenLoop10(i int) int {
	if i >= 1e18 {
		return 19
	}
	x, count := 10, 1
	for x <= i {
		x *= 10
		count++
	}
	return count
}
