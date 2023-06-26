package main

func lengthOfLastWord(s string) int {
	len := len(s)
	lastSpace := 0
	lastChar := 0
	wait := false
	for i := 0; i < len; i++ {
		if s[i] != 32 {
			lastChar = i + 1
			if wait {
				wait = false
				lastSpace = i
			}
		} else {
			wait = true
		}
	}

	return lastChar - lastSpace
}
