package leetcode

func reverseString_2pointers(s string) string {
	array := make([]byte, len(s))
	copy(array, s)
	i := 0
	j := len(s) - 1
	for i < j {
		array[i], array[j] = array[j], array[i]
		i++
		j--
	}
	return string(array)
}
