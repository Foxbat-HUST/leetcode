package medium

func lengthOfLongestSubstring(s string) int {
	sLength := len(s)
	if sLength < 2 {
		return sLength
	}

	charMap := map[byte]int{}

	var start int = 0
	var end int = 1
	var startChar = s[start]
	var endChar = s[end]
	charMap[startChar] = start
	var max int = 1
	for {
		index, exist := charMap[endChar]
		if !exist || index < start {
			charMap[endChar] = end
			tmpMax := end - start + 1
			if tmpMax > max {
				max = tmpMax
			}
			end++
			if end >= sLength {
				break
			}
			endChar = s[end]
			continue
		}

		charMap[endChar] = end

		start = index + 1
		end++
		if end >= sLength || start >= sLength {
			break
		}

		startChar = s[start]
		charMap[startChar] = start

		endChar = s[end]
	}

	return max
}
