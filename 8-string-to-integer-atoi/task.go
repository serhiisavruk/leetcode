package main

func myAtoi(s string) int {
	l := len(s)
	start := 0
	end := 0
	startFound := false
	negative := false
	hasNumber := false
	for i := 0; i < l; i++ {
		if !startFound {
			if s[i] == '-' {
				if i != l-1 && s[i+1] >= 48 && s[i+1] <= 57 {
					negative = true
					continue
				}
			}
			if s[i] >= 48 && s[i] <= 57 {
				start = i
				startFound = true
				hasNumber = true
				continue
			}
		} else {
			if s[i] < 48 || s[i] > 57 {
				end = i
				break
			}
		}
	}
	if !hasNumber {
		return 0
	}
	if hasNumber && end == 0 {
		end = l
	}
	sn := s[start:end]
	l = len(sn)
	n := 0
	for i := 0; i < l; i++ {
		switch sn[i] {
		case '0':
			n += 0
		case '1':
			n += 1
		case '2':
			n += 2
		case '3':
			n += 3
		case '4':
			n += 4
		case '5':
			n += 5
		case '6':
			n += 6
		case '7':
			n += 7
		case '8':
			n += 8
		case '9':
			n += 9
		}
		if i != l-1 {
			n *= 10
		}
	}
	if negative {
		n *= -1
	}
	return n
}
