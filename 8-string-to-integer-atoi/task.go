package main

func myAtoi(s string) int {
	l := len(s)
	start := 0
	end := 0
	startFound := false
	negative := false
	hasNumber := false

	if l > 1 {
		for i := 0; i < l; i++ {
			if !startFound {
				if s[i] == '-' && i < l-1 {
					if s[i+1] < 48 || s[i+1] > 57 {
						break
					}
					negative = true
					continue
				} else if s[i] == '+' {
					if s[i+1] < 48 || s[i+1] > 57 {
						break
					}
				} else if s[i] == ' ' {
					continue
				} else if s[i] >= 48 && s[i] <= 57 {
					start = i
					startFound = true
					hasNumber = true
				} else {
					break
				}
			} else {
				if s[i] < 48 || s[i] > 57 {
					end = i
					break
				}
			}
		}
	} else if l == 1 {
		hasNumber = true
		end = 1
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
		add := 0
		switch sn[i] {
		case '0':
			add = 0
		case '1':
			add = 1
		case '2':
			add = 2
		case '3':
			add = 3
		case '4':
			add = 4
		case '5':
			add = 5
		case '6':
			add = 6
		case '7':
			add = 7
		case '8':
			add = 8
		case '9':
			add = 9
		default:
			add = 0
		}

		if !negative && 2147483647-n < add {
			n = 2147483647
		}

		if negative && 2147483648-n < add {
			n = 2147483648
		}

		n += add

		if !negative && 2147483647-n < add*10 && i != l-1 {
			n = 2147483647
		} else if negative && 2147483648-n < add*10 && i != l-1 {
			n = 2147483648
		} else if i != l-1 {
			n *= 10
		}
	}

	if negative {
		n *= -1
	}
	if n < -2147483648 {
		return -2147483648
	} else if n > 2147483647 {
		return 2147483647
	}
	return n
}
