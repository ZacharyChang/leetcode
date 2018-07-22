package leetcode

func countBits(num int) []int {
	result := make([]int, num+1)
	result[0] = 0
	if num == 0 {
		return result
	}
	result[1] = 1
	// 观察规律
	// 对于 2^n 个元素的数组，后半段数组([2^(n-1):2^n])为前半段数组([0:2^(n-1)])依次拷贝加1
	// 即a[2^(n-1)+m]=a[m]+1,其中 m < 2^n
	for i := 2; i <= num; i++ {
		result[i] = result[getLastPos(i)] + 1
	}
	return result
}

func getLastPos(num int) int {
	i := 1
	for ; i <= num; i *= 2 {

	}
	return num - i/2
}
