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
	for i := 0; i < len(this.queue); i++ {
		if t-this.queue[i] <= 3000 {
			this.queue = this.queue[i:]
			break
		}
	}
	return len(this.queue)
}
