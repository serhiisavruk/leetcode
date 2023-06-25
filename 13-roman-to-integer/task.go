package main

func romanToInt(s string) int {
	nums := map[int32]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}

	len := len(s)

	result := 0
	skipNext := false

	for i, num := range s {
		if skipNext {
			skipNext = false
			continue
		}
		if num == 73 {
			if len-1 > i {
				if s[i+1] == 86 {
					result += 4
					skipNext = true
					continue
				}
				if s[i+1] == 88 {
					result += 9
					skipNext = true
					continue
				}
			}
		} else if num == 88 {
			if len-1 > i {
				if s[i+1] == 76 {
					result += 40
					skipNext = true
					continue
				}
				if s[i+1] == 67 {
					result += 90
					skipNext = true
					continue
				}
			}
		} else if num == 67 {
			if len-1 > i {
				if s[i+1] == 68 {
					result += 400
					skipNext = true
					continue
				}
				if s[i+1] == 77 {
					result += 900
					skipNext = true
					continue
				}
			}
		}

		result += nums[num]
	}

	return result
}
