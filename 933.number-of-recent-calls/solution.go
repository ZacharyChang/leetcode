package leetcode

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	idx := binarySearch(this.queue, t-3000)
	this.queue = this.queue[idx:]
	return len(this.queue)
}

func binarySearch(array []int, target int) int {
	start, end := 0, len(array)-1
	if array[start] >= target {
		return start
	}
	if array[end] <= target {
		return end
	}
	for start < end-1 {
		mid := (start + end) / 2
		if array[mid] == target {
			return mid
		} else if array[mid] > target {
			end = mid
		} else {
			start = mid
		}
	}
	return end
}
