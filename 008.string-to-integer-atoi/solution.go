package leetcode

const INT_MAX = int(int32(^uint32(0) >> 1))
const INT_MIN = int(^INT_MAX)

func myAtoi(str string) int {
	if len(str) == 0 {
		return 0
	}
	if str[0] == ' ' {
		return myAtoi(str[1:])
	}
	if (str[0] >= 'a' && str[0] <= 'z') || (str[0] >= 'A' && str[0] <= 'Z') {
		return 0
	}
	isNega := false
	if str[0] == '-' {
		isNega = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}
	sum := 0
	for i := 0; i < len(str); i++ {
		if str[i] <= '9' && str[i] >= '0' {
			sum = sum*10 + int(str[i]-'0')
			if sum > INT_MAX {
				if isNega {
					return int(INT_MIN)
				}
				return int(INT_MAX)
			}
		} else {
			break
		}
	}
	if isNega {
		return 0 - sum
	}
	return sum
}
