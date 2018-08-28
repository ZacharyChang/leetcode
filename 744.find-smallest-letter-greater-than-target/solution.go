package leetcode

func nextGreatestLetter(letters []byte, target byte) byte {
	start, end := 0, len(letters)-1
	for start < end-1 {
		mid := (start + end) / 2
		if letters[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	if letters[start] > target {
		return letters[start]
	}
	if letters[end] > target {
		return letters[end]
	}
	return letters[0]
}
