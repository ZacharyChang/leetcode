package leetcode

func sortArrayByParity(A []int) []int {
	begin, end := 0, len(A)-1
	for begin < end {
		if A[begin]%2 == 0 {
			begin++
			continue
		}
		if A[end]%2 == 1 {
			end--
			continue
		}
		A[begin], A[end] = A[end], A[begin]
		begin++
		end--
	}
	return A
}
