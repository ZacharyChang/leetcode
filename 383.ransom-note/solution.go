package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	array := make([]int, 'z'-'a'+1)
	// producer
	for i := 0; i < len(magazine); i++ {
		array[magazine[i]-'a']++
	}
	// Consumer
	for i := 0; i < len(ransomNote); i++ {
		array[ransomNote[i]-'a']--
		if array[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
