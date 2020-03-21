package _19_bulb_switcher

func bulbSwitch(n int) int {
	res := 0
	for i := 0; i*i <= n; i++ {
		res = i
	}
	return res
}
