package main

func longestCommonPrefix(strs []string) string {
	if strs[0] == "" {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	index := strs[0][0]
	count := 0
	good := true

	for {
		good = search(strs, count, index)
		if !good {
			break
		}
		count += 1
		if len(strs[0]) == count {
			break
		}
		index = strs[0][count]
	}

	return strs[0][:count]
}

func search(strs []string, count int, index uint8) bool {
	good := true
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) <= count {
			good = false
			break
		}
		if strs[i][count] != index {
			good = false
			break
		}
	}
	return good
}
