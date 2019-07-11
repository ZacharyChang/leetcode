package leetcode

func lemonadeChange(bills []int) bool {
	money := make(map[int]int, 0)
	for _, v := range bills {
		if !change(v, money) {
			return false
		}
	}
	return true
}

func change(bill int, money map[int]int) bool {
	mon := bill - 5
	money[bill]++
	if money[10] >= (mon/10) && money[5] >= (mon%10/5) {
		money[10] -= mon / 10
		money[5] -= mon % 10 / 5
		return true
	}
	if money[5] >= (mon / 5) {
		money[5] -= mon / 5
		return true
	}
	return false
}
